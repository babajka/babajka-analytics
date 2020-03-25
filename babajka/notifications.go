package babajka

import (
	"log"

	"github.com/slack-go/slack"
)

func pushSlackNotification(apiToken, channelName, text string) error {
	api := slack.New(apiToken, slack.OptionDebug(false))
	channelID, timestamp, err := api.PostMessage(channelName, slack.MsgOptionText(text, false))
	if err != nil {
		return err
	}
	log.Printf("Message successfully sent to channel %s at %s\n", channelID, timestamp)
	return nil
}
