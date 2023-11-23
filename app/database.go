package app

import (
	"log"
	"todo-api/models"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var (
	host     = "localhost"
	user     = "postgres"
	password = ""
	dbPort   = "5432"
	dbName   = "todorest"
	db       *gorm.DB
	err      error
	// dns      = "postgres://postgres:postgres@localhost:5432/goreactmovies?sslmode=disable"
)

func StartDB()  {
	dsn := "host=localhost user=postgres password='' dbname=todorest port=5432 sslmode=disable"

	db, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil{
		log.Fatal("error connecting ti database:", err)
	}

	db.Debug().AutoMigrate(models.Todo{})
}

func GetDB() *gorm.DB{
	return db
}