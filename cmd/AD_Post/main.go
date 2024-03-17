package main

import (
	"AD_Post/db"
	"AD_Post/models"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

func main() {
	DDBFileName := "ad.db"
	db.ConnectDB(DDBFileName)

	router := mux.NewRouter()
	router.HandleFunc("/api/v1/ad", list_data).Methods("GET")
	router.HandleFunc("/api/v1/ad", add_data).Methods("POST")

	fmt.Println("Server is running")
	log.Fatal(http.ListenAndServe(":8000", router))
}

func PrintAd(input models.Ad) {
	fmt.Println("title:\t\t", input.Title)
	fmt.Println("StartAt:\t", input.StartAt)
	fmt.Println("EndAt:\t\t", input.EndAt)
	fmt.Println("AgeStart:\t", input.AgeStart)
	fmt.Println("AgeEnd:\t\t", input.AgeEnd)
	fmt.Println("Gender:\t\t", input.Gender)
	fmt.Println("Country:\t", input.Country)
	fmt.Println("Platform:\t", input.Platform)
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

	res, err := db.QueryAd(offset, limit, age, gender, country, platform)
	if err != nil {
		http.Error(w, "Failed to query data from database", http.StatusInternalServerError)
		return
	}

	var ResSlice []models.ResJson
	for _, i := range res {
		ResSlice = append(ResSlice, models.ResJson{Title: i.Title, EndAt: i.EndAt})
	}
	OriJson := struct {
		Item []models.ResJson `json:"item"`
	}{Item: ResSlice}
	ResJson, err := json.Marshal(OriJson)
	if err != nil {
		http.Error(w, "Error marshaling JSON:", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(ResJson)
}

func add_data(w http.ResponseWriter, r *http.Request) {
	// add new data
	var userData models.JsonParse
	err := json.NewDecoder(r.Body).Decode(&userData)
	fmt.Println(r.Body)
	if err != nil {
		http.Error(w, "Failed to parse JSON", http.StatusBadRequest)
		return
	}

	err = db.InsertAd(userData)
	if err != nil {
		http.Error(w, "Failed to add data to database", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	w.Write([]byte("Data added successfully\n"))
}
