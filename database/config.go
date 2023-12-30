package database

import (
  "gorm.io/driver/postgres"
  "gorm.io/gorm"
	"os"

)

var DB *gorm.DB

func InitDB() {
	var err error 
	
	dsn := os.Getenv("DB_URL")

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}
	DB = db
}