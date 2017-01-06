package disturbing

import (
	bot "github.com/kechako/gopher-bot"
	"github.com/kechako/gopher-bot/utils"
)

var keywords = []string{
	"ヤバイ",
	"やばい",
	"やべぇ",
	"やべえ",
	"ヤベェ",
	"ヤベエ",
	"マジか",
	"まじか",
	"すごい",
	"凄い",
	"激しい",
	"心配",
	"危険",
	"危ない",
	"不穏",
	"危機",
}

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

	event.PostMessage("穏やかじゃないわね")

	return true
}

func (p *plugin) Help() string {
	return `disturbing:
	穏やかじゃないわね
`
}
