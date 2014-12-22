package main

import (
	"fmt"
	"net/http"
	"os"
	"santagram_api/server/handler"
)



func main() {
	port := os.Getenv("PORT")
	http.HandleFunc("/users", handler.Handler)
	http.ListenAndServe(fmt.Sprintf(":%v", port), nil)
}
