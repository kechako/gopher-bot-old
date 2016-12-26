package zundoko

import (
	"bytes"
	"math/rand"
	"strings"
	"time"

	bot "github.com/kechako/gopher-bot"
)

const (
	zun     = "ズン"
	doko    = "ドコ"
	kiyoshi = "キ・ヨ・シ！"
)

var (
	keywords = []string{"zundoko", "ズンドコ", "ずんどこ"}
	random   = rand.New(rand.NewSource(time.Now().UnixNano()))
)

type plugin struct {
}

func NewPlugin() bot.Plugin {
	return &plugin{}
}

func (p *plugin) DoAction(event bot.EventInfo) bool {
	if !hasKeywords(event.Text()) {
		return false
	}

	zundoko := [2]string{zun, doko}
	good := [5]string{zun, zun, zun, zun, doko}

	var current [5]string

	reply := bytes.NewBuffer(make([]byte, 0, 1024))

	for current != good {
		shift(&current)
		zd := zundoko[random.Intn(2)]
		current[4] = zd
		reply.WriteString(zd)
	}

	reply.WriteString(kiyoshi)

	event.PostMessage(reply.String())

	return true
}

func hasKeywords(text string) bool {
	for _, k := range keywords {
		if strings.Contains(text, k) {
			return true
		}
	}

	return false
}

func (p *plugin) Help() string {
	return `zundoko:
	ズンドコキヨシ
    `
}

func shift(a *[5]string) {
	a[0], a[1], a[2], a[3], a[4] = a[1], a[2], a[3], a[4], ""
}

var _ bot.Plugin = (*plugin)(nil)
