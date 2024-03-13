package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	router := mux.NewRouter()
	router.HandleFunc("/list/", list_data).Methods("GET")
	router.HandleFunc("/add/", add_data).Methods("POST")

	fmt.Println("Server is running")
	log.Fatal(http.ListenAndServe(":8000", router))
}

func list_data(w http.ResponseWriter, r *http.Request) {
	// list all data
}

func add_data(w http.ResponseWriter, r *http.Request) {
	// add new data
}
