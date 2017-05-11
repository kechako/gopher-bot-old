package japaripark

import (
	"bufio"
	"context"
	"fmt"
	"strings"

	friends "github.com/kechako/go-friends"
	bot "github.com/kechako/gopher-bot"
)

type plugin struct {
	f *friends.Friends
}

// NewPlugin returns a new plugin.
func NewPlugin(appID string) bot.Plugin {
	return &plugin{
		f: friends.New(appID),
	}
}

func (p *plugin) Hello(info bot.BotInfo) {
}

func (p *plugin) DoAction(event bot.EventInfo) bool {
	done := false
	scanner := bufio.NewScanner(strings.NewReader(event.Text()))
	for scanner.Scan() {
		text := scanner.Text()
		if !strings.Contains(text, "得意") {
			continue
		}
		s, err := p.f.Say(context.Background(), text)
		if err != nil {
			event.PostMessage(fmt.Sprintf("Error : %v", err))
			return false
		}
		if s != "" {
			event.PostMessage(s)
			done = true
		}
	}

	return done
}

func (p *plugin) Help() string {
	return `ジャパリパーク:
    きみは Golang が得意なフレンズなんだね
`
}
