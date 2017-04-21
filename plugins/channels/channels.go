package channels

import (
	"bytes"
	"fmt"

	bot "github.com/kechako/gopher-bot"
	"github.com/nlopes/slack"
)

const keyword = "チャンネル一覧"

type plugin struct {
	client *slack.Client
}

// NewPlugin returns a new plugin.
func NewPlugin() bot.Plugin {
	return &plugin{}
}

func (p *plugin) Hello(info bot.BotInfo) {
	p.client = info.GetClient()
}

func (p *plugin) DoAction(event bot.EventInfo) bool {
	if event.Text() != keyword {
		return false
	}

	channels, err := p.client.GetChannels(true)
	if err != nil {
		event.PostMessage(fmt.Sprintf("[Error] Could not get channels : %v", err))
		return true
	}

	var text bytes.Buffer
	for _, ch := range channels {
		fmt.Fprintf(&text, "<#%s|%s> : %s\n", ch.ID, ch.Name, ch.Purpose.Value)
	}

	event.PostMessage(text.String())

	return true
}

func (p *plugin) Help() string {
	return `チャネル一覧: チャンネル一覧を表示します。
    `
}

var _ bot.Plugin = (*plugin)(nil)
