package main

import (
	"AD_Post/db"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"strconv"

	"github.com/gorilla/mux"
)

func main() {
	DDBFileName := "ad.db"
	if _, err := os.Stat(DDBFileName); os.IsNotExist(err) {
		models.BuildDB(DDBFileName)     // create database if the file is not exist
		models.CreateTable(DDBFileName) // create table if the table is not exist
	} else {
		models.ConnectDB(DDBFileName)
	}

	fmt.Println("database and table created")

	router := mux.NewRouter()
	router.HandleFunc("/api/v1/ad", list_data).Methods("GET")
	router.HandleFunc("/api/v1/ad", add_data).Methods("POST")

	fmt.Println("Server is running")
	log.Fatal(http.ListenAndServe(":8000", router))
}

func list_data(w http.ResponseWriter, r *http.Request) {
	// list all data
	QueryParams := r.URL.Query()
	// require condition
	offset, _ := strconv.Atoi(QueryParams.Get("offset"))
	limit, _ := strconv.Atoi(QueryParams.Get("limit"))

	// optional condition
	age := QueryParams.Get("age")
	gender := QueryParams.Get("gender")
	country := QueryParams.Get("country")
	platform := QueryParams.Get("platform")

	res, err := models.QueryAd(offset, limit, age, gender, country, platform)
	if err != nil {
		http.Error(w, "Failed to query data from database", http.StatusInternalServerError)
		return
	}
	for _, i := range res {
		fmt.Println(i.Title)
	}
}

func add_data(w http.ResponseWriter, r *http.Request) {
	// add new data
	var userData models.Ad
	err := json.NewDecoder(r.Body).Decode(&userData)
	if err != nil {
		http.Error(w, "Failed to parse JSON", http.StatusBadRequest)
		return
	}

	err = models.InsertAd(userData)
	if err != nil {
		http.Error(w, "Failed to add data to database", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	w.Write([]byte("Data added successfully\n"))
}
