package userdao

import (
	"gopkg.in/mgo.v2/bson"
	mgo "gopkg.in/mgo.v2"
	"santagram_api/server/user"
	"errors"
)

type MongoDAO struct {
	DialString 	string
	Database  	string
	Collection	string
}

func NewMongoDAO(database string, collection string, dialString string) (MongoDAO) {
	return MongoDAO{DialString: dialString, Database: database, Collection: collection}
}

func (m MongoDAO) Store(storedUser user.User) (*user.User, error) {
	if (storedUser.Username == "") {
		return nil, errors.New("Username is required.")
	}
	if (storedUser.Password == "") {
		return nil, errors.New("Password is required.")
	}
	if (storedUser.Email == "") {
		return nil, errors.New("Email is required.")
	}

	result := new(user.User)
	session, err := mgo.Dial(m.DialString)
	if (err != nil){
		return nil, err
	}
	c := session.DB(m.Database).C(m.Collection)
	err = c.Insert(storedUser)
	err = c.Find(bson.M{"username": storedUser.Username}).One(&result)
	if (err != nil){
		return nil, err
	}
	result.Password = "redacted"
	return result, nil
}

func (m MongoDAO) FindByUsername(username string) (*user.User, error) {
	result := new(user.User)
	session, err := mgo.Dial(m.DialString)
	if (err != nil){
		return nil, err
	}
	c := session.DB(m.Database).C(m.Collection)
	err = c.Find(bson.M{"username": username}).One(&result)
	if (err != nil){
		return nil, err
	}
	return result, nil
}

func (m MongoDAO) Authenticate(username string, password string) (bool, error) {
	authUser, err := m.FindByUsername(username)
	if err != nil {
		return false, err
	}
	if authUser.Password == password {
		return true, nil
	}
	return false, nil
}
