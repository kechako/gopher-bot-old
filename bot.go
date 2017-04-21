package bot

import (
	"bytes"
	"fmt"
	"log"

	"github.com/kechako/gopher-bot/utils"
	"github.com/nlopes/slack"
	"github.com/pkg/errors"
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
				b.handleHello()
			case *slack.ConnectedEvent:
				b.info = ev.Info
				b.logf("Info : %v\n", ev.Info)
				b.logf("Connections : %v\n", ev.ConnectionCount)
			case *slack.DisconnectedEvent:
				b.log("Disconnected")
				if ev.Intentional {
					return
				}
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

func (b *Bot) handleHello() {
	info := BotInfo(b)
	for _, p := range b.plugins {
		p.Hello(info)
	}
}

func (b *Bot) handleMessage(e *slack.MessageEvent) {
	// Ignore myself
	if e.User == b.info.User.ID {
		return
	}

	b.DoActionPlugins(newEvent(b, e))
}

// DoActionPlugins calls DoAction of plugins.
func (b *Bot) DoActionPlugins(event EventInfo) bool {
	if b.showHelp(event, event.Channel()) {
		// shown help
		return true
	}

	for _, p := range b.plugins {
		if done := p.DoAction(event); done {
			return true
		}
	}

	return false
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

// BotID retrieves bot user id.
func (b *Bot) BotID() string {
	return b.info.User.ID
}

// PostMessage posts the text to the channnel.
func (b *Bot) PostMessage(text, channel string) {
	b.rtm.SendMessage(b.rtm.NewOutgoingMessage(text, channel))
}

// PostMessageToThread posts the text to the channnel.
func (b *Bot) PostMessageToThread(text, channel, ts string) {
	msg := b.rtm.NewOutgoingMessage(text, channel)
	msg.ThreadTimestamp = ts
	b.rtm.SendMessage(msg)
}

// ReplyMessage replies the text to the user.
func (b *Bot) ReplyMessage(text, user, channel string) {
	b.PostMessage(fmt.Sprintf("<@%s> %s", user, text), channel)
}

// ReplyMessageToThread replies the text to the user.
func (b *Bot) ReplyMessageToThread(text, user, channel, ts string) {
	b.PostMessageToThread(fmt.Sprintf("<@%s> %s", user, text), channel, ts)
}

// AddReaction adds a reaction to the message.
func (b *Bot) AddReaction(name, channel, timestamp string) error {
	err := b.rtm.AddReaction(name, slack.ItemRef{
		Channel:   channel,
		Timestamp: timestamp,
	})
	if err != nil {
		return errors.Wrap(err, "Failed to add reaction.")
	}

	return nil
}

// GetClient returns *slack.Client.
func (b *Bot) GetClient() *slack.Client {
	return b.client
}

// A BotInfo represents bot information.
type BotInfo interface {
	DoActionPlugins(event EventInfo) bool
	BotID() string
	PostMessage(text, channel string)
	PostMessageToThread(text, channel, ts string)
	ReplyMessage(text, user, channel string)
	AddReaction(name, channel, timestamp string) error
	GetClient() *slack.Client
}
