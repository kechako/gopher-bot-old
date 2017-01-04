package cron

import (
	bot "github.com/kechako/gopher-bot"
)

type eventInfo struct {
	info    bot.BotInfo
	channel string
	text    string
}

func newEventInfo(info bot.BotInfo, channel, text string) bot.EventInfo {
	return &eventInfo{
		info:    info,
		channel: channel,
		text:    text,
	}
}

func (e *eventInfo) Channel() string {
	return e.channel
}

func (e *eventInfo) BotID() string {
	return e.info.BotID()
}

func (e *eventInfo) Text() string {
	return e.text
}

func (e *eventInfo) User() string {
	return e.BotID()
}

func (e *eventInfo) Username() string {
	return ""
}

func (e *eventInfo) ReplyTo() []string {
	return nil
}

func (e *eventInfo) PostMessage(msg string) {
	e.info.PostMessage(msg, e.channel)
}

func (e *eventInfo) ReplyMessage(msg, user string) {
	e.info.ReplyMessage(msg, user, e.channel)
}
