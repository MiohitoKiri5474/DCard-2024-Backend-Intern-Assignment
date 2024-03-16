package main

import (
	"AD_Post/db"
	"log"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func main() {
	sqldb, err := gorm.Open(sqlite.Open("ad.db"), &gorm.Config{})
	if err != nil {
		log.Fatal("Failed to connect to database:", err)
	}
	sqldb.AutoMigrate(&models.Ad{})
}
