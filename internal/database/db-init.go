package database

import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func DBInit(dbUrl string) (*gorm.DB, error) {

	db, err := gorm.Open(postgres.Open(dbUrl), &gorm.Config{})

	return db, err

}
