package models

import (
	"time"

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

type ResJson struct {
	Title string    `json:"title"`
	EndAt time.Time `json:"endAt"`
}
