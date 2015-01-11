package handler

import (
	"fmt"
	"net/http"
	"santagram_api/server/user"
	"santagram_api/server/userdao"
	"encoding/json"
	"io/ioutil"
)
func UserRouter(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		StoreHandler(w, r)
	}
}

func StoreHandler(w http.ResponseWriter, r *http.Request) {
	database := "santagram"
	collection := "santagram"
	dialString := "admin:password@ds063170.mongolab.com:63170/santagram"
	mongoDAO := userdao.NewMongoDAO(database, collection, dialString)
	result := new(user.User)

	body, _ := ioutil.ReadAll(r.Body)
	err := json.Unmarshal(body, result)
	if (err != nil) {
		fmt.Fprintf(w, (fmt.Sprintf(`{"error": %s}`, err.Error())))
	}

	result, err = mongoDAO.Store(*result)
	if (err != nil){
		fmt.Fprintf(w, (fmt.Sprintf(`{"error": %s}`, err.Error())))
	}

	responseBytes, _ := json.Marshal(result)
	w.Write(responseBytes)
}

func AuthenticationHandler(w http.ResponseWriter, r *http.Request) {
	database := "santagram"
	collection := "santagram"
	dialString := "admin:password@ds063170.mongolab.com:63170/santagram"
	mongoDAO := userdao.NewMongoDAO(database, collection, dialString)
	result := new(user.User)

	body, _ := ioutil.ReadAll(r.Body)
	_ = json.Unmarshal(body, result)
	success, err := mongoDAO.Authenticate(result.Username, result.Password)
	if err != nil {
		fmt.Fprintf(w, (fmt.Sprintf(`{"error": %s}`, err.Error())))
	}
	if (success == true) {
		fmt.Fprintf(w, `{"authenticated": true}`)
	} else {
		fmt.Fprintf(w, `{"authenticated": false}`)
	}
}
