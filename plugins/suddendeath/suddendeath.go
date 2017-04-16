package suddendeath

import (
	"bufio"
	"strings"

	bot "github.com/kechako/gopher-bot"
	"github.com/kechako/gopher-bot/utils"
	death "github.com/kechako/suddendeath"
)

type plugin struct {
}

// NewPlugin returns a new plugin.
func NewPlugin() bot.Plugin {
	return &plugin{}
}

func (p *plugin) Hello(info bot.BotInfo) {
}

func (p *plugin) DoAction(event bot.EventInfo) bool {
	s := bufio.NewScanner(strings.NewReader(event.Text()))
	for s.Scan() {
		text := s.Text()
		if utils.HasKeywords(text, "突然の") {
			event.PostMessage(death.Generate(text))
		}
	}

	// 他のプラグインの処理も続行
	return false
}

func (p *plugin) Help() string {
	return `suddendeath: 突然の死

	突然の<something>:

    ＿人人人人人人人人人人＿
    ＞　突然のsomething 　＜
    ￣Y^Y^Y^Y^Y^Y^Y^Y^Y^Y￣
`
}
