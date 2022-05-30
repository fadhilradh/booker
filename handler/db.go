package handler

import (
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB
var err error

func StartDBConnection() {
	dsn := "host=localhost user=postgres password=peni1941 dbname=books port=5432 sslmode=disable TimeZone=Asia/Shanghai"
	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	// db.AutoMigrate(&types.Book{})

	if err != nil {
		log.Fatal(err)
	}

	log.Println("Connection established to Postgres!")
	log.Println(DB) 

	
}
// func AddSampleData() {
	
// }