package handler

import (
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func StartDBConnection() {
	dsn := "host=localhost user=postgres password=peni1941 dbname=books port=5432 sslmode=disable TimeZone=Asia/Shanghai"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatal(err)
	}

	log.Println("Connection established to Postgres!")
	log.Println(db)

}