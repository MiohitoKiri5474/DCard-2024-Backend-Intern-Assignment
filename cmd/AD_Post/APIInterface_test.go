package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"reflect"
	"strings"
	"testing"
	"time"
)

func TestAddData(t *testing.T) {
	mx := http.NewServeMux()
	mx.HandleFunc("/ad", AddData)

	Json := `{"title": "Dcard Intern", "startAt": "2024-1-31T03:00:00.000Z", "endAt": "2024-4-6T03:00:00.000Z", "conditions": {"ageStart": 18, "ageEnd": 24, "gender": ["M", "F"], "country": ["TW", "JP", "US"], "platform": ["ios", "android", "web"]}}`
	reader := strings.NewReader(Json)

	req, err := http.NewRequest(http.MethodPost, "/ad", reader)
	if err != nil {
		t.Error(err.Error())
	}

	w := httptest.NewRecorder()

	mx.ServeHTTP(w, req)

	resp := w.Result()
	if resp.StatusCode != http.StatusOK {
		t.Errorf(w.Body.String())
		t.Errorf("Response code is %v", resp.StatusCode)
	}
}

func BenchmarkAddData(b *testing.B) {
	conditions, _ := json.Marshal(map[string]interface{}{
		"AgeStart": 18,
		"AgeEnd":   24,
		"Gender":   []string{"M", "F"},
		"Country":  []string{"TW", "JP", "US"},
		"Platform": []string{"ios", "android", "web"},
	})
	Json, _ := json.Marshal(map[string]interface{}{
		"Title":      "Dcard Intern",
		"StartAt":    time.Now(),
		"EndAt":      time.Now().AddDate(0, 1, 0),
		"Conditions": conditions,
	})

	for i := 0; i < b.N; i++ {
		http.NewRequest("POST", "/ad", bytes.NewBuffer(Json))
	}
}

func TestListData(t *testing.T) {
	_, err := http.NewRequest("GET", "/ad?offset=0&limit=5", nil)
	if err != nil {
		t.Error(err.Error())
	}
}

func BenchmarkListData(b *testing.B) {
	for i := 0; i < b.N; i++ {
		http.NewRequest("GET", "/ad", nil)
	}
}
