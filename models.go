package models

import "gorm.io/gorm"

type User struct {
	gorm.Model
	AgeStart  int
	AgeEnd    int
	Countries []string
	Platforms []string
}
