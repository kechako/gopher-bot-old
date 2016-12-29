package akari

import (
	"fmt"
	"strings"

	bot "github.com/kechako/gopher-bot"
	"github.com/kechako/gopher-bot/utils"
)

const keyword = "大好き"

type plugin struct {
}

func NewPlugin() bot.Plugin {
	return &plugin{}
}

func (p *plugin) DoAction(event bot.EventInfo) bool {
	text := event.Text()
	if !utils.HasKeywords(text, keyword) {
		return false
	}

	i := strings.LastIndex(text, keyword)
	if i == 0 {
		return false
	}

	likes := text[0:i]
	event.PostMessage(fmt.Sprintf("わぁい%s あかり%s大好き", likes, likes))

	return true
}

func (p *plugin) Help() string {
	return `あかり大好き:
	わぁいうすしお あかりうすしお大好き
`
}
