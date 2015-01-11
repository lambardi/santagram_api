package userdao

import (
	"testing"
	"santagram_api/server/userdao"
	"santagram_api/server/user"
)

func Test_userStoreShouldNotReturnAnError(t *testing.T) {
	database := "santagram"
	collection := "santagram_test"
	dialString := "admin:password@ds063170.mongolab.com:63170/santagram"
	mongoDAO := userdao.NewMongoDAO(database, collection, dialString)

	myUser := new(user.User)
	myUser.Username = "the-guy"
	myUser.Email = "guy@thecouch.com"
	myUser.Password = "something"
	myUser, error := mongoDAO.Store(*myUser)
	if (error != nil){
		t.Fail()
	}
}

func Test_userStoreShouldAssignID(t *testing.T) {
	database := "santagram"
	collection := "santagram"
	dialString := "admin:password@ds063170.mongolab.com:63170/santagram"
	mongoDAO := userdao.NewMongoDAO(database, collection, dialString)

	myUser := new(user.User)
	inputUsername := "the-guy"
	inputEmail := "guy@thecouch.com"
	inputPassword := "IamGod1337!!zor"
	myUser.Username = inputUsername
	myUser.Email = inputEmail
	myUser.Password = inputPassword
	myUser, _ = mongoDAO.Store(*myUser)
	if (myUser.ID.Hex() == ""){
		t.Fail()
	}
}

func Test_userStoreShouldReturnUser(t *testing.T) {
	database := "santagram"
	collection := "santagram"
	dialString := "admin:password@ds063170.mongolab.com:63170/santagram"
	mongoDAO := userdao.NewMongoDAO(database, collection, dialString)

	myUser := new(user.User)
	inputUsername := "the-guy"
	inputEmail := "guy@thecouch.com"
	inputPassword := "IamGod1337!!zor"
	myUser.Username = inputUsername
	myUser.Email = inputEmail
	myUser.Password = inputPassword
	myUser, _ = mongoDAO.Store(*myUser)
	if (myUser.Username != inputUsername) {
		t.Fail()
	}
	if (myUser.Email != inputEmail) {
		t.Fail()
	}
}

func Test_userStoreShouldRequireUsername(t *testing.T) {
	database := "santagram"
	collection := "santagram"
	dialString := "admin:password@ds063170.mongolab.com:63170/santagram"
	mongoDAO := userdao.NewMongoDAO(database, collection, dialString)

	myUser := new(user.User)
	myUser.Email = "guy@thecouch.com"
	myUser.Password = "something"
	myUser, err := mongoDAO.Store(*myUser)
	if (err.Error() != "Username is required.") {
		t.Fail()
	}
}

func Test_userStoreShouldRequirePassword(t *testing.T) {
	database := "santagram"
	collection := "santagram"
	dialString := "admin:password@ds063170.mongolab.com:63170/santagram"
	mongoDAO := userdao.NewMongoDAO(database, collection, dialString)

	myUser := new(user.User)
	myUser.Email = "guy@thecouch.com"
	myUser.Username = "the-guy"
	myUser, err := mongoDAO.Store(*myUser)
	if (err.Error() != "Password is required.") {
		t.Fail()
	}
}

func Test_userStoreShouldRequireEmail(t *testing.T) {
	database := "santagram"
	collection := "santagram"
	dialString := "admin:password@ds063170.mongolab.com:63170/santagram"
	mongoDAO := userdao.NewMongoDAO(database, collection, dialString)

	myUser := new(user.User)
	myUser.Username = "the-guy"
	myUser.Password = "something"
	myUser, err := mongoDAO.Store(*myUser)
	if (err.Error() != "Email is required.") {
		t.Fail()
	}
}

func Test_userStoreShouldNotReturnPassword(t *testing.T) {
	database := "santagram"
	collection := "santagram"
	dialString := "admin:password@ds063170.mongolab.com:63170/santagram"
	mongoDAO := userdao.NewMongoDAO(database, collection, dialString)

	myUser := new(user.User)
	inputUsername := "the-guy"
	inputPassword := "IamGod1337!!zor"
	inputEmail := "guy@thecouch.com"
	myUser.Username = inputUsername
	myUser.Password = inputPassword
	myUser.Email = inputEmail
	myUser, _ = mongoDAO.Store(*myUser)
	if (myUser.Username != inputUsername) {
		t.Fail()
	}
	if (myUser.Email != inputEmail) {
		t.Fail()
	}
	if (myUser.Password != "redacted") {
		t.Fail()
	}
}

func Test_userFindByUsernameShouldReturnUser(t *testing.T) {
	database := "santagram"
	collection := "santagram"
	dialString := "admin:password@ds063170.mongolab.com:63170/santagram"
	mongoDAO := userdao.NewMongoDAO(database, collection, dialString)

	myUser := new(user.User)
	username := "findTester"
	password := "PASSWORD"
	email := "guy@thecouch.com"
	myUser.Username = username
	myUser.Password = password
	myUser.Email = email
	myUser, _ = mongoDAO.Store(*myUser)

	myUser, _ = mongoDAO.FindByUsername("findTester")
	if myUser.Username != username {
		t.Fail()
	}
}

func Test_userAuthenticateShouldReturnTrueIfPasswordsMatch(t *testing.T) {
	database := "santagram"
	collection := "santagram"
	dialString := "admin:password@ds063170.mongolab.com:63170/santagram"
	mongoDAO := userdao.NewMongoDAO(database, collection, dialString)

	myUser := new(user.User)
	username := "passwordTester"
	password := "PASSWORD"
	email := "guy@thecouch.com"
	myUser.Username = username
	myUser.Password = password
	myUser.Email = email
	myUser, _ = mongoDAO.Store(*myUser)
	authenticated, _ := mongoDAO.Authenticate(username, "PASSWORD")
	if authenticated == false {
		t.Fail()
	}
}

func Test_userAuthenticateShouldReturnFalseIfPasswordsMatch(t *testing.T) {
	database := "santagram"
	collection := "santagram"
	dialString := "admin:password@ds063170.mongolab.com:63170/santagram"
	mongoDAO := userdao.NewMongoDAO(database, collection, dialString)

	myUser := new(user.User)
	username := "passwordTester"
	password := "PASSWORD"
	email := "guy@thecouch.com"
	myUser.Username = username
	myUser.Password = password
	myUser.Email = email
	myUser, _ = mongoDAO.Store(*myUser)
	authenticated, _ := mongoDAO.Authenticate(username, "SOMENONPASSWORD")
	if authenticated == true {
		t.Fail()
	}
}
