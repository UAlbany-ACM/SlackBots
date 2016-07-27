package main

import (
	"fmt"
	"os"

	slackbot "github.com/BeepBoopHQ/go-slackbot"
	"github.com/nlopes/slack"
	"golang.org/x/net/context"
)

func startRemindBot() {
	remindBot := slackbot.New(os.Getenv("SLACK_API_TOKEN")) // Enviornment variable must be set
	remindBot.Hear("remindBot (.*)").MessageHandler(remindResponseHandler)
	remindBot.Run()
}

func remindResponseHandler(ctx context.Context, bot *slackbot.Bot, evt *slack.MessageEvent) {
	// evt.Msg.Text is the full string that was entered, so lets parse that through
	// using the remindBot library, something like this:
	// response := remindBot.GetDueDate()
	// 	bot.Reply(evt, response, slackbot.WithTyping)

	fmt.Println(evt.Msg.Text)
	bot.Reply(evt, "This Is My Response", slackbot.WithTyping)
}

func main() {
	startRemindBot()
}
