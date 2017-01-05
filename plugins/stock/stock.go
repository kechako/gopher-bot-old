package stock

import (
	"fmt"
	"strconv"
	"strings"

	bot "github.com/kechako/gopher-bot"
)

const commandName = "stock"

type plugin struct {
}

func NewPlugin() bot.Plugin {
	return &plugin{}
}

func (p *plugin) Hello(info bot.BotInfo) {
}

func (p *plugin) DoAction(event bot.EventInfo) bool {
	params := strings.Fields(event.Text())
	if len(params) != 2 || params[0] != commandName {
		return false
	}

	msg, err := p.getStockMsg(params[1])
	if err != nil {
		msg = fmt.Sprintf("Error : %v", err)
	}
	event.PostMessage(msg)

	return true
}

func (p *plugin) getStockMsg(numStr string) (string, error) {
	n, err := strconv.Atoi(numStr)
	if err != nil {
		return "", err
	}

	s, err := queryStock(n)
	if err != nil {
		return "", err
	}

	return fmt.Sprintf("%s %s\n%s%s %s (%s)", s.Company, s.StockNumber, s.CurrentValue, s.Direction, s.Change, s.PercentChange), nil
}

func (p *plugin) Help() string {
	return `stock:
    stock nnnn
        証券コード nnnn の株価を表示します。
`
}
