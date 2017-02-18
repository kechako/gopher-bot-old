package addgopher

import (
	"strings"

	bot "github.com/kechako/gopher-bot"
	"github.com/kechako/textutil"
	"golang.org/x/text/unicode/norm"
)

var keywords = []string{
	"go",
	"golang",
	"gopher",
}

type plugin struct {
}

func NewPlugin() bot.Plugin {
	return &plugin{}
}

func (p *plugin) Hello(info bot.BotInfo) {
}

func (p *plugin) DoAction(event bot.EventInfo) bool {
	if !checkKeywords(event.Text()) {
		return false
	}

	err := event.AddReaction("gopher")
	if err != nil {
		return false
	}

	// 他のプラグインの処理を続行する
	return false
}

func checkKeywords(text string) bool {
	// Normalize Unicode text with
	text = norm.NFKC.String(text)
	words := textutil.SplitWord(text)

	for _, word := range words {
		word = strings.ToLower(word)
		for _, keyword := range keywords {
			if word == keyword {
				return true
			}
		}
	}

	return false
}

func (p *plugin) Help() string {
	return `addgopher:
    "go" に反応してリアクションします。
`
}
