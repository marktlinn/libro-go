package app

import (
	"fmt"
	"os"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var (
	db *gorm.DB
)

func Connect() {
	user := os.Getenv("DB_USER")
	password := os.Getenv("DB_PASSWORD")
	dbName := os.Getenv("DB_NAME")

	dbString := fmt.Sprintf("%s:%s/%s?charset=utf8&parseTime=True&loc=Local", user, password, dbName)

	conn, err := gorm.Open(mysql.Open(dbString), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	db = conn
}

func GetDB() *gorm.DB {
	return db
}
