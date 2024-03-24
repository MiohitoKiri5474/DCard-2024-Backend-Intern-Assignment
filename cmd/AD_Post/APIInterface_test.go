package main

import (
	"AD_Post/db"
	"AD_Post/models"
	"bytes"
	"encoding/json"
	"log"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func BuildDB() {
	// Create DB
	sqldb, err := gorm.Open(sqlite.Open("ad.db"), &gorm.Config{})
	if err != nil {
		log.Fatal("Failed to connect to database:", err)
	}
	sqldb.AutoMigrate(&models.Ad{})
	db.ConnectDB("ad.db")
}

func CreateJson() []byte {
	// Input data declare
	Json := map[string]interface{}{
		"title":   "Dcard Intern",
		"startAt": time.Now(),
		"endAt":   time.Now().AddDate(0, 1, 0),
		"conditions": map[string]interface{}{
			"ageStart": 18,
			"ageEnd":   24,
			"gender":   []string{"M", "F"},
			"platform": []string{"iOS", "Android", "web"},
			"country":  []string{"TW", "JP", "US"},
		},
	}
	JsonBody, _ := json.Marshal(Json)
	return JsonBody
}

func TestAddData(t *testing.T) {
	BuildDB()
	JsonBody := CreateJson()

	// Create request/recorder and set header
	req, err := http.NewRequest("POST", "/ad", bytes.NewBuffer(JsonBody))
	if err != nil {
		t.Fatal(err)
	}
	req.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()

	// Add data into database
	AddData(w, req)

	// Check result
	if w.Code != http.StatusCreated {
		t.Errorf("expected status %d; got %d", http.StatusOK, w.Code)
		t.Errorf("%s", w.Body)
	}
}

func BenchmarkAddData(b *testing.B) {
	BuildDB()
	JsonBody := CreateJson()
	for i := 0; i < b.N; i++ {
		// Create request/recorder and set header
		req, _ := http.NewRequest("POST", "/ad", bytes.NewBuffer(JsonBody))
		req.Header.Set("Content-Type", "application/json")
		w := httptest.NewRecorder()

		// Add data into database
		AddData(w, req)
	}
}

func TestListData(t *testing.T) {
	BuildDB()

	// Create request and recoder
	req, err := http.NewRequest("GET", "/ad?offset=0&limit=5", nil)
	if err != nil {
		t.Error(err.Error())
	}
	w := httptest.NewRecorder()

	// List data from database with conditions
	ListData(w, req)

	// Check result
	if w.Code != http.StatusOK {
		t.Errorf("expected status %d; got %d", http.StatusOK, w.Code)
		t.Errorf("%s", w.Body)
	}
}

func BenchmarkListData(b *testing.B) {
	BuildDB()
	for i := 0; i < b.N; i++ {
		// Create request and recoder
		req, _ := http.NewRequest("GET", "/ad?offset=0&limit=5", nil)
		w := httptest.NewRecorder()

		// List data from database with conditions
		ListData(w, req)
	}
}
