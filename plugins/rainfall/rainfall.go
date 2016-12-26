package rainfall

import (
	"fmt"
	"io"
	"strings"

	"github.com/kechako/gopher-bot"
)

const (
	RainfallPrefix = "rainfall"
)

type plugin struct {
	appID    string
	locStore *LocationStore
	cmd      *Command
}

type PluginCloser interface {
	bot.Plugin
	io.Closer
}

func NewPlugin(appID string, path string) (PluginCloser, error) {
	locStore := NewLocationStore(path)
	err := locStore.Load()
	if err != nil {
		return nil, err
	}

	p := &plugin{
		appID:    appID,
		locStore: locStore,
	}

	p.cmd = NewCommand(p)

	return p, nil
}

func (p *plugin) Close() error {
	return p.locStore.Save()
}

func (p *plugin) DoAction(event bot.EventInfo) bool {
	params := strings.Fields(event.Text())
	if len(params) == 0 || params[0] != RainfallPrefix {
		return false
	}

	result, err := p.cmd.Execute(params[1:])
	if err != nil {
		if err == CommandSyntaxError {
			event.PostMessage(p.buildHelp())
		} else {
			event.PostMessage(err.Error())
		}
		return true
	}

	event.PostMessage(result)

	return true
}

func (p *plugin) Help() string {
	return `rainfall: 雨チェック
	指定された座標で雨が降っているかどうか表示します。

	rainfall <latitude> <longitude>
	    指定された座標で雨が降っているかどうか表示します。

	rainfall <name>
	    指定された名前の座標で雨が降っているかどうか表示します。

	rainfall add <name> <latitude> <longitude>
	    指定された名前で座標を登録します。

	rainfall change <name> <latitude> <longitude>
	    指定された名前の座標を変更します。

	rainfall rm <name>
	    指定された名前の座標を削除します。

	rainfall list
	    登録された座標を一覧表示します。
    `
}

func (p *plugin) buildHelp() string {
	return fmt.Sprintf("```\n%s\n```", p.Help())
}

var _ bot.Plugin = (*plugin)(nil)
