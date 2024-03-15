package models

import (
	"testing"
	"time"
)

func TestCreate(t *testing.T) {
	BuildDB("123.db")
}

func TestCompressJSON(t *testing.T) {
	OriList := []string{"123", "456", "789"}
	ExpectRes := "123 456 789"
	if ExpectRes != CompressJSON(OriList) {
		t.Error("Value Error")
	}
}

func TestCreateTable(t *testing.T) {
	BuildDB("123.db")
	CreateTable("123.db")
}

func TestInsertAd(t *testing.T) {
	InsertAd(Ad{
		Title:   "Testing Ad",
		StartAt: time.Now(),
		EndAt:   time.Now().AddDate(0, 1, 0),
		Conditions: Conditions{
			AgeStart: 18,
			AgeEnd:   30,
			Gender:   []string{"M", "F"},
			Country:  []string{"TW", "JP"},
			Platform: []string{"Android", "iOS"},
		},
	})
}

func TestConnectDB(t *testing.T) {
	BuildDB("123.db")
	ConnectDB("123.db")
}
