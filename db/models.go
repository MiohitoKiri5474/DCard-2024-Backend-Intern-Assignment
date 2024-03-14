package models

import (
	"database/sql"
	"log"
	"time"

	_ "github.com/mattn/go-sqlite3"
)

type Conditions struct {
	AgeStart int      `json:"ageStart" db:"age_start"`
	AgeEnd   int      `json:"ageEnd" db:"age_end"`
	Gender   []string `json:"gender" db:"gender"`
	Country  []string `json:"country" db:"country"`
	Platform []string `json:"platform" db:"platform"`
}

type Ad struct {
	Title      string     `json:"title" db:"title"`
	StartAt    time.Time  `json:"startAt" db:"start_at"`
	EndAt      time.Time  `json:"endAt" db:"end_at"`
	Conditions Conditions `json:"conditions"`
}

var sqldb *sql.DB

func BuildDB(DDBFileName string) {
	// Build db if it is not exist
	db, err := sql.Open("sqlite3", DDBFileName)
	if err != nil {
		log.Fatal("Failed to connect to database:", err)
	}
	defer db.Close()
}

func CreateTable(DDBFileName string) {
	// Create Table
	_, err := sqldb.Exec(`CREATE TABLE IF NOT EXISTS ads (
        id INTEGER PRIMARY KEY AUTOINCREMENT,
        title TEXT,
        start_at TIMESTAMP,
        end_at TIMESTAMP,
        age_start INTEGER,
        age_end INTEGER,
        country TEXT,
        platform TEXT
    )`)
	if err != nil {
		log.Fatal("Failed to create table:", err)
	}
}

func InsertAd(AdData Ad) error {
	// Create Ad
	_, err := sqldb.Exec("INSERT INTO ads (title, start_at, end_at, age_start, age_end, country, platform) VALUES (?, ?, ?, ?, ?, ?, ?)",
		AdData.Title,
		AdData.StartAt,
		AdData.EndAt,
		AdData.Conditions.AgeStart,
		AdData.Conditions.AgeEnd,
		AdData.Conditions.Country,
		AdData.Conditions.Platform)
	return err
}
