package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"

	"user-profile/dbConnection"
	users "user-profile/handlers"
)



func main() {

	log.Println("Running at Port 5005")
	
	r := mux.NewRouter()

	dbConnection.Connect()
	
	r.HandleFunc("/", func (w http.ResponseWriter, r *http.Request)  {
		fmt.Fprint(w,"Base Api")
	})

	r.HandleFunc("/users", users.GetUsers).Methods("OPTIONS","GET")
	r.HandleFunc("/users/{id}",users.GetUser).Methods("OPTIONS", "GET")
	r.HandleFunc("/users", users.CreateUser).Methods("OPTIONS", "POST")
	r.HandleFunc("/login", users.Login).Methods("OPTIONS", "POST")
	
	http.ListenAndServe(":5005", r)

}



