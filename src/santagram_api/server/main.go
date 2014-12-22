package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	mgo "gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

func handler(w http.ResponseWriter, r *http.Request) {
	session, err := mgo.Dial("admin:password@ds063170.mongolab.com:63170/santagram")
	if (err != nil){
		fmt.Fprintf(err.Error())
	}
	c := session.DB(database).C(collection)
	err = c.Find(bson.M{"id": 2}).One(&result)
	if (err != nil){
		fmt.Fprintf(err.Error())
	}
	fmt.Fprintf(result.Username)
}


func main() {
	database := "santagram"
	collection := "santagram"
	result := new(User)
	fmt.Printf(result.Username)

	session, err := mgo.Dial("admin:password@ds063170.mongolab.com:63170/santagram")
	if (err != nil){
		fmt.Printf("error")
		fmt.Printf(err.Error())
	}
	c := session.DB(database).C(collection)
	err = c.Find(bson.M{"id": 2}).One(&result)
	if (err != nil){
		fmt.Printf("error")
		fmt.Printf(err.Error())
	}
	fmt.Printf("--------------")
	fmt.Printf(result.Username)

	port := os.Getenv("PORT")
	http.HandleFunc("/", handler)
	http.ListenAndServe(fmt.Sprintf(":%v", port), nil)

//	userIn := &User{Username: "Benjamin", Email: "lambardi@gmail.com"}
//	userIn.save()
//	userOut, _ := loadUser("Benjamin")
//	fmt.Println(string(userOut.Email))

}

type User struct {
	ID bson.ObjectId `bson:"_id,omitempty"`
	Username string	 `bson:"username,omitempty"`
	Email string     `bson:"Email,omitempty"`
}

func (u *User) save() error {
	filename := u.Username + ".txt"
	emailBytes := []byte(fmt.Sprintf("%s\n", u.Email))
	return ioutil.WriteFile(filename, emailBytes, 0600)
	}

func loadUser(username string) (*User, error) {
	filename := username + ".txt"
	email, err := ioutil.ReadFile(filename)
	if err != nil {
		return nil, err
	}
	emailString := string(email[:])
	return &User{Username: username, Email: emailString}, nil
}
