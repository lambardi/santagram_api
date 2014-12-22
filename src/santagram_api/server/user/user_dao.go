package user

import "gopkg.in/mgo.v2/bson"

type User struct {
	ID bson.ObjectId `bson:"_id,omitempty"`
	Username string	 `bson:"username,omitempty"`
	Email string     `bson:"Email,omitempty"`
}

func (*User) Save() error {
	return nil
}
