package dice

import (
	"fmt"
	"math/rand"
	"strconv"
	"strings"
	"time"

	bot "github.com/kechako/gopher-bot"
)

const keyword = "dice"

var random = rand.New(rand.NewSource(time.Now().UnixNano()))

type plugin struct {
}

func NewPlugin() bot.Plugin {
	return &plugin{}
}

func (p *plugin) Hello(info bot.BotInfo) {
}

func (p *plugin) DoAction(event bot.EventInfo) bool {
	fields := strings.Fields(event.Text())
	if len(fields) < 1 || fields[0] != keyword {
		return false
	}

	cnt := 1000
	if len(fields) > 1 {
		if n, err := strconv.Atoi(fields[1]); err == nil && n > 0 {
			cnt = n
		}
	}

	event.PostMessage(fmt.Sprint(random.Intn(cnt) + 1))

	return true
}

func (p *plugin) Help() string {
	return `dice:
    dice
        1 〜 1000 の数字をランダムに表示します。

    dice n
        1 〜 n の数字をランダムに表示します。
`
}
