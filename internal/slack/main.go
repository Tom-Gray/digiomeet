package slack

import (
	"DigioChat/internal/matcher"
	"github.com/slack-go/slack"
)

type Client struct {
	client *Client
}

func NewSlackClient(token string) *slack.Client {
	api := slack.New(token, slack.OptionDebug(true))

	return api
}

func (c *Client) getUsers() []matcher.User {
	users, err := c.GetUsers()

	return
}
