package echo

import "github.com/kechako/gopher-bot"

type plugin struct {
}

func NewPlugin() bot.Plugin {
	return &plugin{}
}

func (p *plugin) DoAction(e bot.EventInfo) bool {
	e.PostMessage(e.Text())

	return true
}
