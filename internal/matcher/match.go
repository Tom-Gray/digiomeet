package matcher

import "time"

//User is a slack user who will be matched.
type User struct {
	name string
	id   string
}

//Match is a pairing of two slack users.
type Match struct {
	First  User
	second User
	date   time.Time
}

func makePairing() {}

func recordPairing() {}
