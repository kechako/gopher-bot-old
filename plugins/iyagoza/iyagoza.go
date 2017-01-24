package iyagoza

import (
	"math/rand"
	"time"

	bot "github.com/kechako/gopher-bot"
)

var (
	random   = rand.New(rand.NewSource(time.Now().UnixNano()))
	messages = []*randMsg{
		newRandMsg([]string{
			"いやでござる",
			"なんでや！",
			"お断りします！",
		}),
		newRandMsg([]string{
			"かしこま",
			"かしこま娘",
			"かしこま！",
			"かしこま。",
			"かしこま？",
			"かしこま…",
			"かしこま…。",
			"かしこま〜！",
			"かしこまくら",
			"かしこまって",
			"かしこまっ！",
			"かしこま！！",
			"かしこま～！",
			"かしこま〜！！",
			"かしこまこま！",
			"かしこまだよ！",
			"かしこまった！",
			"かしこまって！",
			"かしこまる〜！",
			"かしこま参上。",
			"かしこま娘め。",
			"かしこまりました",
			"かしこまんじゅう",
			"かしこま〜っ！！",
			"かしこまった〜。",
			"かしこまった～。",
			"かしこま角砂糖！",
			"かしこまりがとう！",
			"かしこまりました。",
			"かしこまバンザイ！",
			"かしこまレシーブ！",
			"かしこまらないクマ。",
			"かしこまパワーです。",
			"かしこまりましたでしょ。",
			"かしこまぷりキュピコーン！",
		}),
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
	event.ReplyMessage(msg.GetMessage(), event.User())
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
