package japaripark

import (
	"fmt"
	"strings"

	bot "github.com/kechako/gopher-bot"
)

var (
	specializedKeywords = []string{
		"が得意",
		"得意",
	}
)

type plugin struct {
}

func NewPlugin() bot.Plugin {
	return &plugin{}
}

func (p *plugin) Hello(info bot.BotInfo) {
}

func (p *plugin) DoAction(event bot.EventInfo) bool {
	text := event.Text()
	i := matchIndex(text)
	if i < 0 {
		return false
	}

	specialized := text[0:i]
	event.PostMessage(fmt.Sprintf("すごーい！きみは%sが得意なフレンズなんだね！", specialized))

	return true
}

func matchIndex(text string) int {
	for _, keyword := range specializedKeywords {
		i := strings.Index(text, keyword)
		if i > 0 {
			return i
		}
	}

	return -1
}

func (p *plugin) Help() string {
	return `ジャパリパーク:
    きみは Golang が得意なフレンズなんだね
`
}
