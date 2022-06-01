package matcher

import (
	"fmt"
	"github.com/Tom-Gray/digiomeet/internal/slackclient"
	"log"
	"math/rand"
	"os"
	"time"
)

var channelID = "C03EK95RQHG" //

//User is a slack user
type User struct {
	name string
	id   string
}

//Users are a lot of users
type Users struct {
	users []User
}

//Match is a pairing of two slack users.
type Match struct {
	first  User
	second User
	date   time.Time
}

type Application struct {
	SlackClient slackclient.SlackClient
}

type SlackClient interface {
	GetAllUserIDs(channelID string) []string
	GetUserNameFromIDs(userID string) (string, error)
	CheckIfBotUser(userID string) (bool, error)
}

func NewApp(slackApp slackclient.SlackClient) *Application {

	app := &Application{
		SlackClient: slackApp,
	}
	return app
}

func Run(app *Application) {

}

func (a *Application) AssembleUsers() {
	userList := a.SlackClient.GetAllUserIDs(channelID)
	userList = a.removeBots(userList)
	a.generateMatches(userList)
	// have users in a big list
	// use something to select two at random
	// remove the two from the list.
	// repeat
	//

	// make match
	fmt.Printf("Number of users: %v", len(userList))

	// check the match.

}

func (a *Application) generateMatches(users []string) {

	userCount := len(users) / 2
	// figure out what to do if there is an uneven number of matches
	if userCount%2 != 0 {
		fmt.Println("uneven matches. What should I do?")
	}
	rand.Seed(time.Now().Unix())
	i := 1
	for i <= userCount {

		fmt.Println(i)
		user := users[0] //select first user from the list

		// loop until we find a valid match for user[0]
		for {
			index := rand.Intn(len(users))
			if index == 0 {
				index++ // Can't allow the user to be matched with themselves at position 0
			}
			match := users[index] //get a match
			if !validateMatch() {
				continue
			}

			fmt.Printf("user : %v\nmatched with: %v\n", user, match) //make some kind of data structure

			// add pairing to structure
			// serialise to a file.
			users = removeFromSliceByIndex(users, index) //remove matched person from slice
			users = removeFromSliceByIndex(users, 0)
			break
		}

		i++

	}

	os.Exit(0)

}
func (a *Application) getUsernameFromID(users []string) {
	for _, userID := range users {
		name, err := a.SlackClient.GetUserNameFromIDs(userID)
		if err != nil {
			fmt.Printf("User %v dropped.", name)
		}
		//u := User{name: name, id: userID}
		fmt.Printf("User ID: %v:%v\n", userID, name)
	}

}

func (a *Application) ValidatePairing() bool {
	return true
}

func recordPairing() {

	fmt.Println("Recording Pair")
}

func (a *Application) removeBots(users []string) []string {
	var botUsers []string
	for _, userID := range users {
		isBot, err := a.SlackClient.CheckIfBotUser(userID)
		if err != nil {
			log.Println(err)
			continue
		}
		if isBot {
			botUsers = append(botUsers, userID)
		}

	}

	for _, bot := range botUsers {
		users = remove(users, bot)
	}

	return users

}

func remove(s []string, r string) []string {
	for i, v := range s {
		if v == r {
			return append(s[:i], s[i+1:]...)
		}
	}
	return s
}

//validateMatch checks to see if the users have been paired before.
func validateMatch() bool {
	return true
}

func removeFromSliceByIndex(s []string, index int) []string {

	s[index] = s[len(s)-1] // copy last element toindex we wish to remove
	s[len(s)-1] = ""       // remove the element
	s = s[:len(s)-1]

	return s
}
