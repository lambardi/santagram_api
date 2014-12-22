package user

import (
	"testing"
	"santagram_api/server/user"
)

func Test_userSaveShouldNotReturnAnError(t *testing.T) {
	myUser := new(user.User)
	myUser.Username = "the-guy"
	myUser.Email = "guy@thecouch.com"
	error := myUser.Save()
	if (error != nil){
		t.Fail()
	}
}

func Test_userSaveShouldAssignID(t *testing.T) {
	myUser := new(user.User)
	myUser.Username = "the-guy"
	myUser.Email = "guy@thecouch.com"
	myUser.Save()
	if (myUser.ID.Hex() == ""){
		t.Fail()
	}
}
