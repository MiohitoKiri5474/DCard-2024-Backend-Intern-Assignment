package main

import (
	"AD_Post/db"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	DDBFileName := "ad.db"
	db.ConnectDB(DDBFileName)

	router := mux.NewRouter()
	router.HandleFunc("/ad", list_data).Methods("GET")
	router.HandleFunc("/ad", add_data).Methods("POST")

	fmt.Println("Server is running")
	log.Fatal(http.ListenAndServe(":8000", router))
}
