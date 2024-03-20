package main

import (
	"AD_Post/db"
	"AD_Post/models"
	"encoding/json"
	"fmt"
	"net/http"
	"reflect"
	"strconv"
)

func ListData(w http.ResponseWriter, r *http.Request) {
	// list all data
	defer func() {
		if r := recover(); r != nil {
			errMsg := fmt.Sprintf("Invalid parameter: %v", r)
			http.Error(w, errMsg, http.StatusBadRequest)
		}
	}()
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
	if err := CheckAge(age); err != nil {
		panic(err.Error())
	}
	gender := QueryParams.Get("gender")
	if err := CheckGender(gender); err != nil {
		panic(err.Error())
	}
	country := QueryParams.Get("country")
	if err := CheckCountry(country); err != nil {
		panic(err.Error())
	}
	platform := QueryParams.Get("platform")
	if err := CheckPlatform(platform); err != nil {
		panic(err.Error())
	}

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

func AddData(w http.ResponseWriter, r *http.Request) {
	// add new data
	defer func() {
		if r := recover(); r != nil {
			errMsg := fmt.Sprintf("Invalid parameter: %v", r)
			http.Error(w, errMsg, http.StatusBadRequest)
		}
	}()

	var userData models.JsonParse
	fmt.Println(reflect.TypeOf(r.Body))
	if err := json.NewDecoder(r.Body).Decode(&userData); err != nil {
		http.Error(w, "Failed to parse JSON", http.StatusBadRequest)
		return
	}

	if err := CheckJSon(userData); err != nil {
		panic(err.Error())
	}

	if err := db.InsertAd(userData); err != nil {
		http.Error(w, "Failed to add data to database", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	w.Write([]byte("Data added successfully\n"))
}
