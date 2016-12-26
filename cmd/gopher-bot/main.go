package main

import (
	"flag"
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/kechako/gopher-bot"
	"github.com/kechako/gopher-bot/plugins/echo"
)

var slackToken string

func init() {
	flag.StringVar(&slackToken, "token", os.Getenv("SLACK_TOKEN"), "Slack API token.")
}

func main() {
	flag.Parse()

	bot := bot.New(slackToken)
	bot.SetLogger(log.New(os.Stdout, "slack-bot: ", log.LstdFlags))

	bot.AddPlugin(echo.NewPlugin())

	sig := make(chan os.Signal)
	signal.Notify(sig, syscall.SIGINT, syscall.SIGTERM, syscall.SIGQUIT)
	go func() {
		<-sig
		bot.Stop()
	}()
	bot.Start()
}
