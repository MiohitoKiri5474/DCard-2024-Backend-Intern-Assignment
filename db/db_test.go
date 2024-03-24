package db

import (
	"AD_Post/models"
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
	InputJSON := models.JsonParse{
		Title:   "Testing Ad",
		StartAt: time.Now(),
		EndAt:   time.Now().AddDate(0, 1, 0),
		Conditions: models.Conditions{
			AgeStart: 18,
			AgeEnd:   30,
			Gender:   []string{"M", "F"},
			Country:  []string{"TW", "JP"},
			Platform: []string{"Android", "iOS"},
		},
	}
	res := ConvertToAd(InputJSON)
	cmp := models.Ad{
		Title:    "Testing Ad",
		StartAt:  InputJSON.StartAt,
		EndAt:    InputJSON.EndAt,
		AgeStart: 18,
		AgeEnd:   30,
		Gender:   "M F ",
		Country:  "TW JP ",
		Platform: "Android iOS ",
	}

	if res != cmp {
		t.Error("Value Error")
	}
}

func TestConnectDB(t *testing.T) {
	ConnectDB("123.db")
	sqldb.AutoMigrate(&models.Ad{})
}

func TestInsertAd(t *testing.T) {
	InsertAd(models.JsonParse{
		Title:   "Testing Ad",
		StartAt: time.Now(),
		EndAt:   time.Now().AddDate(0, 1, 0),
		Conditions: models.Conditions{
			AgeStart: 18,
			AgeEnd:   30,
			Gender:   []string{"M", "F"},
			Country:  []string{"TW", "JP"},
			Platform: []string{"Android", "iOS"},
		},
	})
}

func BenchmarkInsertAd(b *testing.B) {
	for i := 0; i < b.N; i++ {
		InsertAd(models.JsonParse{
			Title:   "Testing Ad",
			StartAt: time.Now(),
			EndAt:   time.Now().AddDate(0, 1, 0),
			Conditions: models.Conditions{
				AgeStart: 18,
				AgeEnd:   30,
				Gender:   []string{"M", "F"},
				Country:  []string{"TW", "JP"},
				Platform: []string{"Android", "iOS"},
			},
		})
	}
}

func TestQueryAd(t *testing.T) {
	QueryAd(15, 20, "24", "M", "TW", "iOS")
}

func BenchmarkQueryAd(b *testing.B) {
	for i := 0; i < b.N; i++ {
		QueryAd(15, 20, "24", "M", "TW", "iOS")
	}
}
