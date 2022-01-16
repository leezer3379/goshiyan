package common

import (
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"goshiyan/model"
)
var (
	DB *gorm.DB
)

func InitDB() *gorm.DB {
	// github.com/mattn/go-sqlite3
	db, err := gorm.Open(sqlite.Open("gorm.db"), &gorm.Config{})
	if err != nil {
		panic(err)
	}

	db.AutoMigrate(&model.User{})

	DB = db
	return db
}

func GetDB() *gorm.DB {
	return DB
}

