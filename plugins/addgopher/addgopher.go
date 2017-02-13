package addgopher

import (
	"strings"

	bot "github.com/kechako/gopher-bot"
)

const keyword = "go"

type plugin struct {
}

func NewPlugin() bot.Plugin {
	return &plugin{}
}

func (p *plugin) Hello(info bot.BotInfo) {
}

func (p *plugin) DoAction(event bot.EventInfo) bool {
	if !strings.Contains(strings.ToLower(event.Text()), keyword) {
		return false
	}

	err := event.AddReaction("gopher")
	if err != nil {
		return false
	}

	// 他のプラグインの処理を続行する
	return false
}

func (p *plugin) Help() string {
	return `addgopher:
    "go" に反応してリアクションします。
`
}
