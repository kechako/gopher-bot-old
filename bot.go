package bot

import (
	"bytes"
	"fmt"
	"log"

	"github.com/kechako/gopher-bot/utils"
	"github.com/nlopes/slack"
)

// A Bot represents a bot client.
type Bot struct {
	client *slack.Client
	rtm    *slack.RTM
	info   *slack.Info

	plugins []Plugin

	logger *log.Logger
}

// New creates a new Bot.
func New(token string) *Bot {
	return &Bot{
		client: slack.New(token),
	}
}

// AddPlugin adds the plugin p to the bot.
func (b *Bot) AddPlugin(p Plugin) {
	b.plugins = append(b.plugins, p)
}

// Start starts the bot.
func (b *Bot) Start() {
	rtm := b.client.NewRTM()
	b.rtm = rtm

	go rtm.ManageConnection()

	for {
		select {
		case msg := <-rtm.IncomingEvents:
			switch ev := msg.Data.(type) {
			case *slack.HelloEvent:
				b.log("Hello")
			case *slack.ConnectedEvent:
				b.info = ev.Info
				b.logf("Info : %v\n", ev.Info)
				b.logf("Connections : %v\n", ev.ConnectionCount)
			case *slack.DisconnectedEvent:
				b.log("Disconnected")
				return
			case *slack.MessageEvent:
				b.handleMessage(ev)
			}
		}
	}
}

// Stop stops the bot.
func (b *Bot) Stop() {
	b.rtm.Disconnect()
}

func (b *Bot) handleMessage(e *slack.MessageEvent) {
	// Ignore myself
	if e.User == b.info.User.ID {
		return
	}

	event := newEvent(b, e)

	if b.showHelp(event, e.Channel) {
		// shown help
		return
	}

	for _, p := range b.plugins {
		if done := p.DoAction(event); done {
			break
		}
	}
}

func (b *Bot) showHelp(event EventInfo, channel string) bool {
	if !(utils.IsReplyToBot(event.BotID(), event.ReplyTo()) && utils.HasKeywords(event.Text(), "help")) {
		return false
	}

	var buf bytes.Buffer

	buf.WriteString("```\n")

	for _, p := range b.plugins {
		buf.WriteString(p.Help())
		buf.WriteString("\n")
	}

	buf.WriteString("```")

	b.PostMessage(buf.String(), channel)

	return true
}

// PostMessage posts the text to the channnel.
func (b *Bot) PostMessage(text, channel string) {
	b.rtm.SendMessage(b.rtm.NewOutgoingMessage(text, channel))
}

// ReplyMessage replies the text to the user.
func (b *Bot) ReplyMessage(text, user, channel string) {
	b.PostMessage(fmt.Sprintf("<@%s> %s", user, text), channel)
}
