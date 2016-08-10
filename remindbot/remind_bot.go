package remindbot

import (
	"errors"
	"os"
	"strings"

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
	referencedAssignment, err := parseAssignment(evt.Msg.Text)
	if err != nil {
		bot.Reply(evt, "Unrecognized Pattern", slackbot.WithTyping)
	}
	remindText := bot.GetDueDate(referencedAssignment)
	bot.Reply(evt, "Here is your reminder: "+remindText, slackbot.WithTyping)
}

func parseAssignment(fullInput string) (string, error) {
	params := strings.Split(fullInput, " ")
	if len(params) > 2 {
		return " ", errors.New("Unrecognized Pattern")
	}
	return params[1], nil
}
