package main

import (
	"github.com/Tom-Gray/digiomeet/internal/matcher"
	"github.com/Tom-Gray/digiomeet/internal/slackclient"
	"os"
)

func main() {
	token := os.Getenv("SLACK_TOKEN")
	slackApp := slackclient.NewSlackClient(token)
	app := matcher.NewApp(slackApp)
	app.AssembleUsers()

	//match
}
