package iyagoza

import bot "github.com/kechako/gopher-bot"

type plugin struct {
}

func NewPlugin() bot.Plugin {
	return &plugin{}
}

func (p *plugin) DoAction(event bot.EventInfo) bool {
	if !isReplyToBot(event.BotID(), event.ReplyTo()) {
		return false
	}

	event.ReplyMessage("いやでござる", event.User())
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

var _ bot.Plugin = (*plugin)(nil)
