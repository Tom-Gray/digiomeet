package slackclient

import (
	"fmt"
	"github.com/slack-go/slack"
	"log"
)

type SlackClient struct {
	client *slack.Client
}

func NewSlackClient(token string) SlackClient {
	c := slack.New(token, slack.OptionDebug(true))

	return SlackClient{c}

}

func (c *SlackClient) GetAllUserIDs(channelID string) []string {
	var opts = &slack.GetUsersInConversationParameters{
		ChannelID: channelID,
		Cursor:    "",
		Limit:     0,
	}

	usersInChannel, _, err := c.client.GetUsersInConversation(opts)
	if err != nil {
		log.Fatalln(err)
	}

	fmt.Printf("discovered %v users\n", len(usersInChannel))
	for _, user := range usersInChannel {
		fmt.Printf("Found user %v\n", user)
	}

	return usersInChannel
}

func (c *SlackClient) GetUserNameFromIDs(userID string) (string, error) {
	userInfo, err := c.client.GetUserInfo(userID)
	if err != nil {
		return "", err
	}

	return userInfo.Name, nil
}

// CheckIfBotUser Returns true of the user is a bot.
func (c *SlackClient) CheckIfBotUser(userID string) (bool, error) {
	userInfo, err := c.client.GetUserInfo(userID)
	if err != nil {
		return false, err
	}
	if userInfo.IsBot == true {
		return true, nil
	} else {
		return false, nil
	}

}
