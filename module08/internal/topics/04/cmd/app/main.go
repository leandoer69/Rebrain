package main

import (
	"Rebrain/module08/internal/topics/04/internal/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func main() {
	dsn := "host=localhost user=db_user password=pwd123 dbname=gorm port=54320 sslmode=disable"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}

	err = models.InitModels(db)
	if err != nil {
		panic(err)
	}

	Start(db)
}
