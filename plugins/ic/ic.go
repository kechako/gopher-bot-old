package ic

import (
	"math/rand"
	"time"

	bot "github.com/kechako/gopher-bot"
	"github.com/kechako/gopher-bot/utils"
)

const keyword = "なるほど"

var feedbacks = []string{
	"なるほど?",
	"なるほど!",
	"なるほど!!",
	"なるほど!!!",
	"な〜るほど!",
	"なるほど〜なるほど!",
	"なるほどなるほどですぞ!",
}

type plugin struct {
	r *rand.Rand
}

func NewPlugin() bot.Plugin {
	return &plugin{
		r: rand.New(rand.NewSource(time.Now().UnixNano())),
	}
}

func (p *plugin) Hello(info bot.BotInfo) {
}

func (p *plugin) DoAction(event bot.EventInfo) bool {
	if !utils.HasKeywords(event.Text(), keyword) {
		return false
	}

	event.PostMessage(feedbacks[p.r.Intn(len(feedbacks))])

	return true
}

func (p *plugin) Help() string {
	return `なるほど:
	「なるほど」に反応して相槌を打ちます。
`
}
