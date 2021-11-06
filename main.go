package main

import (
	"github.com/gorilla/mux"
	"net/http"
	"log"
)


func initializeRouter() {
	r := mux.NewRouter()

	r.HandleFunc("/users",GetUsers).Methods("GET")
	r.HandleFunc("/users/{id}",GetUser).Methods("GET")
	r.HandleFunc("/users",CreateUser).Methods("POST")
	r.HandleFunc("/users/{id}",UpdateUser).Methods("PUT")
	r.HandleFunc("/users/{id}",DeleteUsers).Methods("DELETE")

	http.ListenAndServe(":5002", r)

	log.Println("Server is up at port 5002")
}

func main() {
	InitializeMigration()
	initializeRouter()
}