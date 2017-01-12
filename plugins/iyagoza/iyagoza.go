package iyagoza

import (
	"math/rand"
	"time"

	bot "github.com/kechako/gopher-bot"
)

var (
	random   = rand.New(rand.NewSource(time.Now().UnixNano()))
	messages = []string{
		"いやでござる",
		"かしこまっ！",
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
	if !isReplyToBot(event.BotID(), event.ReplyTo()) {
		return false
	}

	msg := messages[random.Intn(len(messages))]
	event.ReplyMessage(msg, event.User())
	return true
}

func isReplyToBot(botID string, replyTo []string) bool {
	for _, r := range replyTo {
		if r == botID {
			return true
		}
	}

	return false
}

func (p *plugin) Help() string {
	return `iyagoza:
	reply 'いやでござる'
    `
}

var _ bot.Plugin = (*plugin)(nil)
