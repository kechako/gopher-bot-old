package disturbing

import (
	"math/rand"
	"time"

	bot "github.com/kechako/gopher-bot"
	"github.com/kechako/gopher-bot/utils"
)

var keywords = []string{
	"ヤバイ",
	"やばい",
	"やべぇ",
	"やべえ",
	"ヤベェ",
	"ヤベエ",
	"マジか",
	"まじか",
	"すごい",
	"凄い",
	"激しい",
	"心配",
	"危険",
	"危ない",
	"不穏",
	"危機",
}

var messages = []string{
	"穏やかじゃ…。",
	"穏やかじゃない。",
	"穏やかじゃない！",
	"穏やかじゃない？",
	"穏やかじゃない…。",
	"穏やかじゃないし",
	"穏やかじゃないって",
	"穏やかじゃないね。",
	"穏やかじゃないの？",
	"穏やかじゃないわ。",
	"穏やかじゃないわ！",
	"穏やかじゃない気持",
	"穏やかじゃない話！",
	"穏やかじゃないうえに",
	"穏やかじゃないかも。",
	"穏やかじゃないけど",
	"穏やかじゃないことが",
	"穏やかじゃないって。",
	"穏やかじゃないはず。",
	"穏やかじゃないわね。",
	"穏やかじゃない予感。",
	"穏やかじゃなかった。",
	"穏やかじゃなさすぎて",
	"穏やかじゃありません。",
	"穏やかじゃない　だね。",
	"穏やかじゃないことが",
	"穏やかじゃないっぽい。",
	"穏やかじゃないと思う。",
	"穏やかじゃない高さね。",
	"穏やかじゃなさすぎる。",
	"穏やかじゃなさすぎる！",
	"穏やかじゃナッシング！",
	"穏やかじゃない案件だね。",
	"穏やかじゃない状況！！",
	"穏やかじゃない！　本物！",
	"穏やかじゃないだろうから。",
	"穏やかじゃないニュース。",
	"穏やかじゃない！　でしょ？",
	"穏やかじゃないに決まってる。",
	"穏やかじゃないわね　らいち。",
	"穏やかじゃないどころじゃない。",
	"穏やかじゃない　スーパーレア写真！",
}

var random = rand.New(rand.NewSource(time.Now().UnixNano()))

type plugin struct {
}

func NewPlugin() bot.Plugin {
	return &plugin{}
}

func (p *plugin) Hello(info bot.BotInfo) {
}

func (p *plugin) DoAction(event bot.EventInfo) bool {
	if !utils.HasKeywords(event.Text(), keywords...) {
		return false
	}

	event.PostMessage(messages[random.Intn(len(messages))])

	return true
}

func (p *plugin) Help() string {
	return `disturbing:
	穏やかじゃないわね
`
}
