package japaripark

import (
	"fmt"
	"strings"

	"github.com/ikawaha/kagome/tokenizer"
	bot "github.com/kechako/gopher-bot"
)

var (
	specializedKeywords = []string{
		"が得意",
		"得意",
	}
)

type plugin struct {
	t tokenizer.Tokenizer
}

// NewPlugin returns a new plugin.
func NewPlugin() bot.Plugin {
	return &plugin{
		t: tokenizer.New(),
	}
}

func (p *plugin) Hello(info bot.BotInfo) {
}

func (p *plugin) DoAction(event bot.EventInfo) bool {
	s, ok := p.getSpeciality(event.Text())
	if !ok {
		return false
	}

	event.PostMessage(fmt.Sprintf("すごーい！きみは%sが得意なフレンズなんだね！", s))

	return true
}

func (p *plugin) getSpeciality(text string) (string, bool) {
	tokens := p.t.Tokenize(text)
	if len(tokens) == 0 {
		return "", false
	}

	l := len(tokens)
	words := make([]string, 0, l)
	hasSpeciality := false
	for i := 0; i < l; i++ {
		t := tokens[i]

		if t.Class == tokenizer.DUMMY {
			continue
		}

		var next tokenizer.Token
		if i+1 < l {
			next = tokens[i+1]
		}

		if matchKeywordToken(t, next) {
			hasSpeciality = true
		} else if matchParticle(t, "が") {
			var moreNext tokenizer.Token
			if i+2 < l {
				moreNext = tokens[i+2]
			}
			if matchKeywordToken(next, moreNext) {
				hasSpeciality = true
			}
		}
		if hasSpeciality {
			break
		}

		words = append(words, t.Surface)
	}

	if !hasSpeciality || len(words) == 0 {
		return "", false
	}

	return strings.Join(words, ""), true
}

func matchParticle(t tokenizer.Token, p string) bool {
	return t.Surface == p && t.Features()[0] == "助詞"
}

func matchKeywordToken(t, next tokenizer.Token) bool {
	// token matches "得意"
	if t.Surface == "得意" && t.Features()[0] == "名詞" {
		// next token does not match "先"
		if next.Features()[0] == "名詞" {
			// not "得意"
			return false
		}

		return true
	}

	return false
}

func (p *plugin) Help() string {
	return `ジャパリパーク:
    きみは Golang が得意なフレンズなんだね
`
}
