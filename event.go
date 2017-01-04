package bot

import (
	"regexp"

	"github.com/nlopes/slack"
)

type event struct {
	bot   *Bot
	event *slack.MessageEvent
}

// A EventInfo represents an event of a received message.
type EventInfo interface {
	Channel() string
	BotID() string
	Text() string
	User() string
	Username() string
	ReplyTo() []string
	PostMessage(string)
	ReplyMessage(string, string)
}

// Create a new event struct.
func newEvent(b *Bot, ev *slack.MessageEvent) *event {
	return &event{
		bot:   b,
		event: ev,
	}
}

// Channel retrieves the channnel.
func (e *event) Channel() string {
	return e.event.Channel
}

// BotID retrieves the ID of the bot.
func (e *event) BotID() string {
	return e.bot.info.User.ID
}

// Text retrieves the text of the received message.
func (e *event) Text() string {
	return e.event.Text
}

// User retrieves the ID of user who sent the message.
func (e *event) User() string {
	return e.event.User
}

// Username retrieves the name of user who sent the message.
func (e *event) Username() string {
	u := e.bot.info.GetUserByID(e.event.User)
	if u == nil {
		return ""
	}
	return u.Name
}

var replyRegex = regexp.MustCompile("<@[0-9a-zA-Z]+>")

// ReplyTo retrieves a slice of the userID of the reply.
func (e *event) ReplyTo() []string {
	users := replyRegex.FindAllString(e.event.Text, -1)

	for i := 0; i < len(users); i++ {
		u := users[i]
		users[i] = u[2 : len(u)-1]
	}

	return users
}

// PostMessage posts the text to the channnel.
func (e *event) PostMessage(msg string) {
	e.bot.PostMessage(msg, e.event.Channel)
}

// ReplyMessage replies the text to the user.
func (e *event) ReplyMessage(msg, user string) {
	e.bot.ReplyMessage(msg, user, e.event.Channel)
}
