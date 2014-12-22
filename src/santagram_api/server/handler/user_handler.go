package handler

import (
	mgo "gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
	"fmt"
	"net/http"
	"santagram_api/server/user"
)
func Handler(w http.ResponseWriter, r *http.Request) {
	database := "santagram"
	collection := "santagram"
	result := new(user.User)

	session, err := mgo.Dial("admin:password@ds063170.mongolab.com:63170/santagram")
	if (err != nil){
		fmt.Fprintf(w, err.Error())
	}
	c := session.DB(database).C(collection)
	err = c.Find(bson.M{"id": 2}).One(&result)
	if (err != nil){
		fmt.Fprintf(w, err.Error())
	}
	fmt.Fprintf(w, result.ID.Hex())
}
