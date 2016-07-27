package remindbot

import (
	"os"

	"golang.org/x/net/context"

	"github.com/BeepBoopHQ/go-slackbot"
	"github.com/nlopes/slack"
)

func startRemindBot() {
	remindBot := slackbot.New(os.Getenv("SLACK_API_TOKEN")) // Enviornment variable must be set
	remindBot.Hear("remindBot (.*)").MessageHandler(remindResponseHandler)
	remindBot.Run()
}

func remindResponseHandler(ctx context.Context, bot *slackbot.Bot, evt *slack.MessageEvent) {
	// evt.Msg.Text is the full string that was entered, so lets parse that through
	// using the remindBot library, something like this:
	// remindText := remindBot.GetDueDate()
	// 	bot.Reply(evt, response, slackbot.WithTyping)

	bot.Reply(evt, "Here is your reminder: "+remindText, slackbot.WithTyping)
}
