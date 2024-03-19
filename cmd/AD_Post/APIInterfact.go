package main

import (
	"AD_Post/db"
	"AD_Post/models"
	"encoding/json"
	"net/http"
	"strconv"
)

func list_data(w http.ResponseWriter, r *http.Request) {
	// list all data
	QueryParams := r.URL.Query()
	// require condition
	OffsetStr, LimitStr := QueryParams.Get("offset"), QueryParams.Get("limit")
	offset, limit := 5, 5
	if OffsetStr != "" {
		offset, _ = strconv.Atoi(OffsetStr)
	}
	if LimitStr != "" {
		limit, _ = strconv.Atoi(LimitStr)
	}

	// optional condition
	age := QueryParams.Get("age")
	if !AgeCheck(age) {
		http.Error(w, "Invalid age parameter: out of range", http.StatusBadRequest)
	}
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
