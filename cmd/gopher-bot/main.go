package main

import (
	"flag"
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/kechako/gopher-bot"
	"github.com/kechako/gopher-bot/plugins/addgopher"
	"github.com/kechako/gopher-bot/plugins/akari"
	"github.com/kechako/gopher-bot/plugins/channels"
	"github.com/kechako/gopher-bot/plugins/cron"
	"github.com/kechako/gopher-bot/plugins/dice"
	"github.com/kechako/gopher-bot/plugins/disturbing"
	"github.com/kechako/gopher-bot/plugins/ic"
	"github.com/kechako/gopher-bot/plugins/iyagoza"
	"github.com/kechako/gopher-bot/plugins/japaripark"
	"github.com/kechako/gopher-bot/plugins/lgtm"
	"github.com/kechako/gopher-bot/plugins/ppap"
	"github.com/kechako/gopher-bot/plugins/rainfall"
	"github.com/kechako/gopher-bot/plugins/stock"
	"github.com/kechako/gopher-bot/plugins/suddendeath"
	"github.com/kechako/gopher-bot/plugins/zundoko"
	//"github.com/kechako/gopher-bot/plugins/echo"
)

var (
	slackToken     string
	yahooAppID     string
	rainfallPath   string
	cronPath       string
	disturbingPath string
)

func init() {
	flag.StringVar(&slackToken, "token", os.Getenv("SLACK_TOKEN"), "Slack API token.")
	flag.StringVar(&yahooAppID, "appid", os.Getenv("YAHOO_APP_ID"), "Yahoo App Id.")
	flag.StringVar(&rainfallPath, "rainfall-path", os.Getenv("RAINFALL_PATH"), "Rainfall plugin data store path.")
	flag.StringVar(&cronPath, "cron-path", os.Getenv("CRON_PATH"), "Cron plugin data store path.")
	flag.StringVar(&disturbingPath, "disturbing-path", os.Getenv("DISTURBING_PATH"), "Disturbing plugin config path.")
}

func main() {
	flag.Parse()

	bot := bot.New(slackToken)
	bot.SetLogger(log.New(os.Stdout, "slack-bot: ", log.LstdFlags))

	//bot.AddPlugin(echo.NewPlugin())

	c, err := cron.NewPlugin(cronPath)
	if err != nil {
		panic(err)
	}
	defer c.Close()

	rain, err := rainfall.NewPlugin(yahooAppID, rainfallPath)
	if err != nil {
		panic(err)
	}
	defer rain.Close()

	dist, err := disturbing.NewPlugin(disturbingPath)
	if err != nil {
		panic(err)
	}

	bot.AddPlugin(c)
	bot.AddPlugin(rain)
	bot.AddPlugin(channels.NewPlugin())
	bot.AddPlugin(addgopher.NewPlugin())
	bot.AddPlugin(dice.NewPlugin())
	bot.AddPlugin(stock.NewPlugin())
	bot.AddPlugin(suddendeath.NewPlugin())
	bot.AddPlugin(japaripark.NewPlugin())
	bot.AddPlugin(akari.NewPlugin())
	bot.AddPlugin(ic.NewPlugin())
	bot.AddPlugin(dist)
	bot.AddPlugin(zundoko.NewPlugin())
	bot.AddPlugin(ppap.NewPlugin())
	bot.AddPlugin(lgtm.NewPlugin())
	bot.AddPlugin(iyagoza.NewPlugin())

	sig := make(chan os.Signal)
	signal.Notify(sig, syscall.SIGINT, syscall.SIGTERM, syscall.SIGQUIT)
	go func() {
		<-sig
		bot.Stop()
	}()
	bot.Start()
}
