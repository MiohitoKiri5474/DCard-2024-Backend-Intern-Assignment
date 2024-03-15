package models

import (
	"database/sql"
	"fmt"
	"log"
	"strconv"
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
	ID         int
	Title      string     `json:"title" db:"title"`
	StartAt    time.Time  `json:"startAt" db:"start_at"`
	EndAt      time.Time  `json:"endAt" db:"end_at"`
	Conditions Conditions `json:"conditions"`
}

var sqldb *sql.DB

func CompressJSON(OriList []string) string {
	var res string
	for _, i := range OriList {
		if res != "" {
			res += " "
		}
		res += i
	}
	fmt.Println("\t" + res)
	return res
}

func BuildDB(DDBFileName string) {
	// Build db if it is not exist
	var err error
	sqldb, err = sql.Open("sqlite3", DDBFileName)
	if err != nil {
		log.Fatal("Failed to connect to database:", err)
	}
	fmt.Println("Database Created")
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
	fmt.Println("Table Created")
}

func InsertAd(AdData Ad) error {
	// Create Ad
	fmt.Println(AdData)
	_, err := sqldb.Exec("INSERT INTO ads (title, start_at, end_at, age_start, age_end, country, platform) VALUES (?, ?, ?, ?, ?, ?, ?)",
		AdData.Title,
		AdData.StartAt,
		AdData.EndAt,
		AdData.Conditions.AgeStart,
		AdData.Conditions.AgeEnd,
		CompressJSON(AdData.Conditions.Country),
		CompressJSON(AdData.Conditions.Platform))
	return err
}

func QueryAd(offset int, limit int, age string, gender string, country string, platform string) ([]Ad, error) {
	QueryString := "SELECT * FROM ads WHERE start_at <= ? AND end_at >= ?"
	filter := []interface{}{time.Now(), time.Now()}
	if age != "" {
		QueryString += " AND age_start <= " + age + " AND age_end >= " + age
		ageInt, _ := strconv.Atoi(age)
		filter = append(filter, ageInt, ageInt)
	}
	if gender != "" {
		QueryString += " AND gender = ?"
		filter = append(filter, gender)
	}
	if country != "" {
		QueryString += " AND country = ?"
		filter = append(filter, country)
	}
	if platform != "" {
		QueryString += " AND platform = ?"
		filter = append(filter, platform)
	}
	QueryString += " ORDER BY id ASC LIMIT ? OFFSET ?"
	filter = append(filter, limit, offset)

	rows, err := sqldb.Query(QueryString, filter...)
	if err != nil {
		return []Ad{}, err
	}

	defer rows.Close()

	var res []Ad

	for rows.Next() {
		var ad Ad
		err := rows.Scan(&ad.ID, &ad.Title, &ad.StartAt, &ad.EndAt, &ad.Conditions.AgeStart, &ad.Conditions.AgeEnd, &ad.Conditions.Country, &ad.Conditions.Platform)
		if err != nil {
			return []Ad{}, err
		}
		res = append(res, ad)
	}

	return res, nil
}

func ConnectDB(DDBFileName string) {
	var err error
	sqldb, err = sql.Open("sqlite3", DDBFileName)
	if err != nil {
		log.Fatal("Failed to connect to database:", err)
	}
}
