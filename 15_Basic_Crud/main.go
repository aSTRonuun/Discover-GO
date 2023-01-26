package main

import (
	"fmt"
	"log"
	"net/http"

	"basic_crud/db"
	"basic_crud/server"

	"github.com/gorilla/mux"
)

func main() {
	error := db.CreateSchema()
	if error != nil {
		log.Fatal(error)
	}

	router := mux.NewRouter()
	router.HandleFunc("/users", server.CreateUser).Methods(http.MethodPost)
	router.HandleFunc("/users", server.GetAllUsers).Methods(http.MethodGet)
	router.HandleFunc("/users/{id}", server.GetUserById).Methods(http.MethodGet)
	router.HandleFunc("/users/{id}", server.UpdateUser).Methods(http.MethodPut)
	router.HandleFunc("/users/{id}", server.DeleteUser).Methods(http.MethodDelete)

	fmt.Println("Server is listening on port 5000")
	log.Fatal(http.ListenAndServe(":5000", router))
}
