package models

import (
	"fmt"
	"testing"
	"time"
)

func TestCompressJSON(t *testing.T) {
	OriList := []string{"123", "456", "789"}
	ExpectRes := "123 456 789 "
	if ExpectRes != CompressJSON(OriList) {
		t.Error("Value Error")
	}
}

func TestConvertToAd(t *testing.T) {
	InputJSON := JsonParse{
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
	}
	res := ConvertToAd(InputJSON)
	cmp := Ad{
		Title:    "Testing Ad",
		StartAt:  InputJSON.StartAt,
		EndAt:    InputJSON.EndAt,
		AgeStart: 18,
		AgeEnd:   30,
		Gender:   "M F ",
		Country:  "TW JP ",
		Platform: "Android iOS ",
	}

	fmt.Println("res: ", res)
	fmt.Println("cmp: ", cmp)
	if res != cmp {
		t.Error("Value Error")
	}
}

func TestConnectDB(t *testing.T) {
	ConnectDB("123.db")
	sqldb.AutoMigrate(&Ad{})
}

func TestInsertAd(t *testing.T) {
	InsertAd(JsonParse{
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
