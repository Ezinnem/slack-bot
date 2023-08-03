package main

import (
	"fmt"
	"log"
	"context"
	"os"
	"github.com/shomali11/slacker"
)
func printCommandEvents()

func main(){
	os.Setenv("SLACK_BOT_TOKEN", "xoxb-5688014620353-5675363985907-XLQWDF6621H4rLnEcFGlexSA")
	os.Setenv("SLACK_APP_TOKEN", "xapp-1-A05LK3TA7R6-5699132078592-96f473cf7045ba69e11adbc2ba353555bd3e33e03a65480c40e186dff6cb026e")

	bot := slacker.NewClient(os.Getenv("SLACK_BOT_TOKEN"), os.Getenv("SLACK_APP_TOKEN"))

	go printCommandEvents(bot.CommandEvents())
}

/home/martina/slackbot/main.go