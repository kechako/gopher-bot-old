package bot

import (
	"fmt"
	"log"

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
	for _, p := range b.plugins {
		if done := p.DoAction(event); done {
			break
		}
	}
}

// SetLogger sets the logger l to the bot.
func (b *Bot) SetLogger(l *log.Logger) {
	b.logger = l
}

func (b *Bot) logf(format string, v ...interface{}) {
	if b.logger != nil {
		b.logger.Printf(format, v...)
	}
}

func (b *Bot) log(v ...interface{}) {
	if b.logger != nil {
		b.logger.Print(v...)
	}
}

// PostMessage posts the text to the channnel.
func (b *Bot) PostMessage(text, channel string) {
	b.rtm.SendMessage(b.rtm.NewOutgoingMessage(text, channel))
}

// ReplyMessage replies the text to the user.
func (b *Bot) ReplyMessage(text, user, channel string) {
	b.PostMessage(fmt.Sprintf("<@%s> %s", user, text), channel)
}