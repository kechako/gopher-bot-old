package cron

import (
	"bytes"
	"errors"
	"fmt"
	"strconv"
	"strings"

	bot "github.com/kechako/gopher-bot"
	"github.com/kechako/gopher-bot/plugins/cron/internal/data"
	crn "github.com/robfig/cron"
)

const cronCommandName = "cron"

var errCommandCanNotParse = errors.New("command can not parse.")

const (
	cmdNameList   = "list"
	cmdNameRemove = "rm"
)

type PluginCloser interface {
	Close() error
	bot.Plugin
}

type plugin struct {
	info bot.BotInfo
	cron *crn.Cron

	path  string
	store *data.ScheduleStore
}

func NewPlugin(path string) (PluginCloser, error) {
	store, err := data.LoadScheduleStore(path)
	if err != nil {
		return nil, err
	}

	p := &plugin{
		path:  path,
		store: store,
	}

	p.initCron()

	return p, nil
}

func (p *plugin) Close() error {
	p.cron.Stop()
	return data.SaveScheduleStore(p.path, p.store)
}

func (p *plugin) initCron() {
	p.cron = crn.New()

	for _, s := range p.store.List() {
		p.addCronSchedule(s)
	}
}

func (p *plugin) resetCron() {
	if p.cron != nil {
		p.cron.Stop()
	}

	p.initCron()

	p.cron.Start()
}

func (p *plugin) Hello(info bot.BotInfo) {
	p.info = info
	p.cron.Start()
}

func (p *plugin) DoAction(event bot.EventInfo) bool {
	params := strings.Fields(event.Text())
	if len(params) < 1 || params[0] != cronCommandName {
		return false
	}
	params = params[1:]

	var cmdName string
	if len(params) == 0 {
		cmdName = cmdNameList
	} else {
		cmdName = params[0]
	}

	var msg string
	var err error
	switch cmdName {
	case cmdNameList:
		// list cron
		msg = p.listSchedules()
	case cmdNameRemove:
		msg, err = p.removeSchedule(params[1:])
	default:
		msg, err = p.addSchedule(params, event.Channel())
	}

	if err != nil {
		msg = p.getErrMessage(err)
	}

	event.PostMessage(msg)

	return true
}

func (p *plugin) listSchedules() string {
	var buf bytes.Buffer

	for i, s := range p.store.List() {
		fmt.Fprintf(&buf, "[%d] %s %s\n", i+1, s.Fields, s.Command)
	}

	return buf.String()
}

func (p *plugin) removeSchedule(params []string) (string, error) {
	if len(params) != 1 {
		return "", errCommandCanNotParse
	}

	i, err := strconv.Atoi(params[0])
	if err != nil {
		return "", errCommandCanNotParse
	}

	removed := p.store.Remove(i - 1)
	if !removed {
		return "", fmt.Errorf("No.%d is out of range.", i)
	}

	p.resetCron()

	return fmt.Sprintf("Remove No.%d schedule.", i), nil
}

func (p *plugin) addSchedule(params []string, channel string) (string, error) {
	if len(params) < 2 {
		return "", errCommandCanNotParse
	}

	if params[0][0] != '@' && len(params) < 6 {
		return "", errCommandCanNotParse
	}

	s := &data.Schedule{
		Channel: channel,
	}
	if params[0][0] == '@' {
		s.Fields = params[0]
		s.Command = strings.Join(params[1:], " ")
	} else {
		s.Fields = strings.Join(params[:5], " ")
		s.Command = strings.Join(params[5:], " ")
	}

	err := p.addCronSchedule(s)
	if err != nil {
		return "", err
	}

	p.store.Add(s)

	return "Schedule added.", nil
}

func (p *plugin) addCronSchedule(s *data.Schedule) error {
	cs, err := crn.ParseStandard(s.Fields)
	if err != nil {
		return err
	}

	p.cron.Schedule(cs, p.newCronJob(s))

	return nil
}

func (p *plugin) newCronJob(s *data.Schedule) crn.Job {
	return &cronTask{
		info: p.info,
		s:    s,
	}
}

func (p *plugin) Help() string {
	return `cron:
    cron
        登録されているエントリーをリスト表示。

    cron 0 * * * * * zundoko
        毎時0分に zundoko を実行。

    cron @monthly
        毎月1日0時に1回実行。

    cron rm n
        n番目のエントリーを削除。
`
}

func (p *plugin) buildHelp() string {
	return fmt.Sprintf("```\n%s\n```", p.Help())
}

func (p *plugin) getErrMessage(err error) string {
	if err == errCommandCanNotParse {
		return p.buildHelp()
	} else {
		return fmt.Sprintf("Error : %v", err)
	}
}

type cronTask struct {
	info bot.BotInfo
	s    *data.Schedule
}

func (t *cronTask) Run() {
	done := t.info.DoActionPlugins(newEventInfo(t.info, t.s.Channel, t.s.Command))
	if !done {
		t.info.PostMessage(t.s.Command, t.s.Channel)
	}
}
