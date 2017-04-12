package disturbing

import (
	"io/ioutil"
	"math/rand"
	"regexp"
	"time"

	bot "github.com/kechako/gopher-bot"
	toml "github.com/pelletier/go-toml"
)

type config struct {
	Keywords []string `toml:"keywords"`
	Messages []string `toml:"messages"`
}

func loadConf(name string) (*config, error) {
	buf, err := ioutil.ReadFile(name)
	if err != nil {
		return nil, err
	}

	conf := new(config)
	err = toml.Unmarshal(buf, conf)
	if err != nil {
		return nil, err
	}

	return conf, nil
}

type plugin struct {
	conf     *config
	keywords []*regexp.Regexp
	messages []string
	random   *rand.Rand
}

// NewPlugin returns a new plugin.
func NewPlugin(name string) (bot.Plugin, error) {
	conf, err := loadConf(name)
	if err != nil {
		return nil, err
	}

	keywords := make([]*regexp.Regexp, len(conf.Keywords))

	for i, k := range conf.Keywords {
		r, err := regexp.Compile(k)
		if err != nil {
			return nil, err
		}
		keywords[i] = r
	}

	return &plugin{
		conf:     conf,
		keywords: keywords,
		messages: conf.Messages,
		random:   rand.New(rand.NewSource(time.Now().UnixNano())),
	}, nil
}

func (p *plugin) Hello(info bot.BotInfo) {
}

func (p *plugin) DoAction(event bot.EventInfo) bool {
	if !p.matchKeywords(event.Text()) {
		return false
	}

	event.PostMessage(p.messages[p.random.Intn(len(p.messages))])

	return true
}

func (p *plugin) matchKeywords(text string) bool {
	for _, keyword := range p.keywords {
		if keyword.MatchString(text) {
			return true
		}
	}

	return false
}

func (p *plugin) Help() string {
	return `disturbing:
	穏やかじゃないわね
`
}
