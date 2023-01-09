package database

import (
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"

	"smart-home-backend/model"
)

var Instance *gorm.DB
var err error

func Connect() {
	InstanceURL := ""
	Instance, err = gorm.Open(postgres.Open(InstanceURL), &gorm.Config{})
	if err != nil {
		log.Fatal(err)
		panic("Cannot connect to DB")
	}
	log.Println("Connected to Database...")
}

func Migrate() {

	//AutoMigrate will create tables, missing foreign keys, constraints, columns and indexes
	Instance.AutoMigrate(&model.Request{})
	Instance.AutoMigrate(&model.Command{})
	log.Println("Database Migration Completed...")
}
