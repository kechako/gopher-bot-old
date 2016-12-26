package main

import (
	"flag"
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/kechako/gopher-bot"
	"github.com/kechako/gopher-bot/plugins/iyagoza"
	"github.com/kechako/gopher-bot/plugins/rainfall"
	"github.com/kechako/gopher-bot/plugins/zundoko"
	//"github.com/kechako/gopher-bot/plugins/echo"
)

var (
	slackToken   string
	yahooAppId   string
	rainfallPath string
)

func init() {
	flag.StringVar(&slackToken, "token", os.Getenv("SLACK_TOKEN"), "Slack API token.")
	flag.StringVar(&yahooAppId, "appid", os.Getenv("YAHOO_APP_ID"), "Yahoo App Id.")
	flag.StringVar(&rainfallPath, "rainfall-path", os.Getenv("RAINFALL_PATH"), "Rainfall plugin data store path.")
}

func main() {
	flag.Parse()

	bot := bot.New(slackToken)
	bot.SetLogger(log.New(os.Stdout, "slack-bot: ", log.LstdFlags))

	//bot.AddPlugin(echo.NewPlugin())
	rain, err := rainfall.NewPlugin(yahooAppId, rainfallPath)
	if err != nil {
		panic(err)
	}
	rain.Close()

	bot.AddPlugin(rain)
	bot.AddPlugin(zundoko.NewPlugin())
	bot.AddPlugin(iyagoza.NewPlugin())

	sig := make(chan os.Signal)
	signal.Notify(sig, syscall.SIGINT, syscall.SIGTERM, syscall.SIGQUIT)
	go func() {
		<-sig
		bot.Stop()
	}()
	bot.Start()
}
