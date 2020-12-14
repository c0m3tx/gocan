package main

import (
	"fmt"

	"github.com/c0m3tx/gocan"
)

type User struct {
	Name  string
	Level int
}

func (u User) Abilities() gocan.Ability {
	// By default, no ability is set
	var abilities gocan.Ability

	// All user can read "target"
	abilities.Grant(gocan.Read, "target", nil)

	// A simple user comparison function, verifies that the names match
	userCmp := func(user, target interface{}) bool {
		userU, userOk := user.(User)
		targetU, targetOk := target.(User)
		if !userOk || !targetOk {
			return false
		}

		return userU.Name == targetU.Name
	}

	// User can update his own profile
	abilities.Grant(gocan.Update, User{Name: u.Name}, userCmp)

	if u.Level > 5 {
		// high level users can do anything on "target"
		abilities.Grant(gocan.Manage, "target", nil)
	}

	return abilities
}

func main() {
	user := User{Name: "User", Level: 1}
	otherUser := User{Name: "Other user", Level: 5}
	if gocan.Can(user, gocan.Read, "target") {
		fmt.Println("User can read target")
	}

	// This check passes, as user can update its profile...
	if gocan.Can(user, gocan.Update, user) {
		fmt.Println("User can update its profile")
	}
	// ...but this one doesn't!
	if gocan.Can(user, gocan.Update, otherUser) {
		fmt.Println("User can update other user profile")
	}

	// this only succeeds if user level is > 5
	if gocan.Can(user, gocan.Create, "target") {
		fmt.Println("User can create a new target")
	}
}
