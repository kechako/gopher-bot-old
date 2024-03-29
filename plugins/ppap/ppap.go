package ppap

import (
	"bytes"
	"math/rand"
	"time"

	bot "github.com/kechako/gopher-bot"
	"github.com/kechako/gopher-bot/utils"
)

const (
	pen        = "\u2712\ufe0f"
	pineapple  = "\U0001f34d"
	apple      = "\U0001F34E"
	ppapFinish = "ペンパイナッポーアッポーペン"
)

var (
	keywords = []string{"ppap", "PPAP"}
	random   = rand.New(rand.NewSource(time.Now().UnixNano()))
)

type plugin struct {
}

func NewPlugin() bot.Plugin {
	return &plugin{}
}

func (p *plugin) Hello(info bot.BotInfo) {
}

func (p *plugin) DoAction(event bot.EventInfo) bool {
	if !utils.HasKeywords(event.Text(), keywords...) {
		return false
	}

	ppapWords := [3]string{pen, pineapple, apple}
	good := [4]string{pen, pineapple, apple, pen}

	var current [4]string

	reply := bytes.NewBuffer(make([]byte, 0, 1024))

	for current != good {
		shift(&current)
		pa := ppapWords[random.Intn(3)]
		current[3] = pa
		reply.WriteString(pa)
	}

	reply.WriteString(ppapFinish)

	event.PostMessage(reply.String())

	return true
}

func (p *plugin) Help() string {
	return `PPAP:
	` + "\u2712\ufe0f\U0001f34d\U0001F34E\u2712\ufe0f" + `
    `
}

func shift(a *[4]string) {
	a[0], a[1], a[2], a[3] = a[1], a[2], a[3], ""
}

var _ bot.Plugin = (*plugin)(nil)
