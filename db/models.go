package models

import (
	"fmt"
	"log"
	"strings"
	"time"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

type Conditions struct {
	AgeStart int      `json:"ageStart"`
	AgeEnd   int      `json:"ageEnd"`
	Gender   []string `json:"gender"`
	Country  []string `json:"country"`
	Platform []string `json:"platform"`
}

type JsonParse struct {
	Title      string     `json:"title"`
	StartAt    time.Time  `json:"startAt"`
	EndAt      time.Time  `json:"endAt"`
	Conditions Conditions `json:"conditions"`
}

type Ad struct {
	gorm.Model
	Title    string    `json:"title" db:"title"`
	StartAt  time.Time `json:"startAt" db:"start_at"`
	EndAt    time.Time `json:"endAt" db:"end_at"`
	AgeStart int       `json:"ageStart" db:"age_start"`
	AgeEnd   int       `json:"ageEnd" db:"age_end"`
	Gender   string    `json:"gender" db:"gender"`
	Country  string    `json:"country" db:"country"`
	Platform string    `json:"platform" db:"platform"`
}

var sqldb *gorm.DB

func CompressJSON(OriList []string) string {
	var res string
	for _, i := range OriList {
		res += i + " "
	}
	return res
}

func ConvertToAd(input JsonParse) Ad {
	return Ad{
		Title:    input.Title,
		StartAt:  input.StartAt,
		EndAt:    input.EndAt,
		AgeStart: input.Conditions.AgeStart,
		AgeEnd:   input.Conditions.AgeEnd,
		Gender:   CompressJSON(input.Conditions.Gender),
		Country:  CompressJSON(input.Conditions.Country),
		Platform: CompressJSON(input.Conditions.Platform),
	}
}

func ConnectDB(DDBFileName string) {
	// Build db if it is not exist
	var err error
	sqldb, err = gorm.Open(sqlite.Open(DDBFileName), &gorm.Config{})
	if err != nil {
		log.Fatal("Failed to connect to database:", err)
	}
	fmt.Println("Connect Created")
}

func InsertAd(AdData JsonParse) error {
	// Create Ad
	Converted := ConvertToAd(AdData)
	sqldb.Create(&Converted)
	return nil
}

func QueryAd(offset int, limit int, age string, gender string, country string, platform string) ([]Ad, error) {
	// Query Ads from the db
	var res []Ad
	query := sqldb.Model(&Ad{})

	if age != "" {
		query = query.Where("age_start <= ?", age)
		query = query.Where("age_end >= ?", age)
	}

	if country != "" {
		CountryStr := strings.Join(strings.Split(country, ","), " ")
		query = query.Where("country LIKE ?", "%"+CountryStr+"%")
	}

	if platform != "" {
		PlatformStr := strings.Join(strings.Split(platform, ","), " ")
		query = query.Where("platform LIKE ?", "%"+PlatformStr+"%")
	}
	if gender != "" {
		GenderStr := strings.Join(strings.Split(gender, ","), " ")
		query = query.Where("gender LIKE ?", "%"+GenderStr+"%")
	}

	query = query.Order("end_at asc")
	if offset >= 0 && limit > 0 {
		query = query.Offset(offset).Limit(limit)
	}

	if err := query.Find(&res).Error; err != nil {
		return nil, err
	}
	return res, nil
}
