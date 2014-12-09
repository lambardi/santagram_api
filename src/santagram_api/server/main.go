package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
)

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hi there, I love %s!", r.URL.Path[1:])
}


func main() {
	port := os.Getenv("SG_PORT")
	http.HandleFunc("/", handler)
	http.ListenAndServe(fmt.Sprintf(":%v", port), nil)

//	userIn := &User{Username: "Benjamin", Email: "lambardi@gmail.com"}
//	userIn.save()
//	userOut, _ := loadUser("Benjamin")
//	fmt.Println(string(userOut.Email))

}

type User struct {
	Username string
	Email string
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
