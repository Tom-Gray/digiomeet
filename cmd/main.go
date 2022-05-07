package main

import (
	slack2 "DigioChat/internal/slack"
	"fmt"

	"github.com/slack-go/slack"
)

const token = "xoxb-46190970084-3478178351015-ffYKaV7pOD8FLzK9elxouXNh"

func main() {
	api := slack.New(token, slack.OptionDebug(true))
	// If you set debugging, it will log all requests to the console
	// Useful when encountering issues
	// slack.New("YOUR_TOKEN_HERE", slack.OptionDebug(true))
	groups, err := api.GetUserGroups()
	if err != nil {
		fmt.Printf("%s\n", err)
		return
	}
	for _, group := range groups {
		fmt.Printf("ID: %s, Name: %s\n", group.ID, group.Name)
	}

	users, err := api.GetUsers()

	for _, user := range users {
		fmt.Printf("id: %s, Name: %s\n", user.ID, user.Name)
	}

	var convoParams = &slack.GetConversationsParameters{}
	channels, _, err := api.GetConversations(convoParams)

	for _, channel := range channels {
		fmt.Printf("id: %s, Name: %s\n", channel.ID, channel.Name)
	}

	//attachment := slack.Attachment{
	//	Pretext: "some pretext",
	//	Text:    "some text",
	//	// Uncomment the following part to send a field too
	//	/*
	//		Fields: []slack.AttachmentField{
	//			slack.AttachmentField{
	//				Title: "a",
	//				Value: "no",
	//			},
	//		},
	//	*/
	//}

	channelID, timestamp, err := api.PostMessage(
		"C1C5LUM0U",
		slack.MsgOptionText("Hello", false),
	)
	if err != nil {
		fmt.Printf("%s\n", err)
		return
	}
	fmt.Printf("Message successfully sent to channel %s at %s", channelID, timestamp)
}

func run() {
	client := slack2.NewSlackClient(token)
}
