package gocan_test

import (
	"fmt"
	"testing"

	"github.com/c0m3tx/gocan"
)

type User struct {
	Name string
}

func (u User) Abilities() gocan.Ability {
	var abilities gocan.Ability
	abilities.Grant(gocan.Read, "target", nil)

	abilities.Grant(gocan.Manage, "manage target", nil)
	abilities.Deny(gocan.Update, "manage target", nil)

	cmp := func(x interface{}, y interface{}) bool {
		xU, xOk := x.(User)
		yU, yOk := y.(User)

		if !xOk || !yOk {
			return false
		}

		return xU.Name == yU.Name
	}
	abilities.Grant("update profile", User{Name: u.Name}, cmp)
	fmt.Println(abilities)

	return abilities
}

func TestCan(t *testing.T) {
	u := User{Name: "can read target"}
	if !gocan.Can(u, gocan.Read, "target") {
		t.Errorf("Should be able to read target")
	}
	if gocan.Can(u, gocan.Create, "target") {
		t.Errorf("Shouldn't be able to create target")
	}
}

func TestCanManage(t *testing.T) {
	u := User{Name: "can manage target"}
	if !gocan.Can(u, gocan.Read, "manage target") {
		t.Errorf("Unable to read from target")
	}
	if !gocan.Can(u, gocan.Create, "manage target") {
		t.Errorf("Unable to create target")
	}
	if !gocan.Can(u, gocan.Destroy, "manage target") {
		t.Errorf("Unable to destroy target")
	}
	if gocan.Can(u, gocan.Update, "manage target") {
		t.Errorf("Shouldn't be able to update target (denied after granting manage)")
	}
}

func TestCanWithFunction(t *testing.T) {
	u := User{Name: "can update profile"}
	ou := User{Name: "other user"}

	if !gocan.Can(u, "update profile", u) {
		t.Errorf("Should be able to update its own profile")
	}
	if gocan.Can(u, "update profile", ou) {
		t.Errorf("Shouldn't be able to update another profile")
	}
}
