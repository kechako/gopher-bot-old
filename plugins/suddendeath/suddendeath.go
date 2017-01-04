package suddendeath

import (
	bot "github.com/kechako/gopher-bot"
	"github.com/kechako/gopher-bot/utils"
	death "github.com/kechako/suddendeath"
)

type plugin struct {
}

func NewPlugin() bot.Plugin {
	return &plugin{}
}

func (p *plugin) Hello(info bot.BotInfo) {
}

func (p *plugin) DoAction(event bot.EventInfo) bool {
	if !utils.HasKeywords(event.Text(), "突然の") {
		return false
	}

	event.PostMessage(death.Generate(event.Text()))

	return true
}

func (p *plugin) Help() string {
	return `suddendeath: 突然の死

	突然の<something>:

    ＿人人人人人人人人人人＿
    ＞　突然のsomething 　＜
    ￣Y^Y^Y^Y^Y^Y^Y^Y^Y^Y￣
`
}
