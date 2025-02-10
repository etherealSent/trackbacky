package database

import (
	"fmt"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
)

var DB *gorm.DB

func DB_conn() {
	dataSourceName := "host=localhost user=postgres password=p4ssw0rD dbname=tracks_db port=5432 sslmode=disable TimeZone=Europe/Moscow"

	db, err := gorm.Open(postgres.Open(dataSourceName), &gorm.Config{})
	if err != nil {
		log.Fatal(err)
	}
	DB = db
	fmt.Println("Connected to database")
}
