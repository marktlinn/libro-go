package config

import (
	"fmt"
	"os"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type DB struct {
	db *gorm.DB
}

// Creates a connection to the database and returns the DB instance.
func Connect() (*DB, error) {
	user := os.Getenv("DB_USER")
	password := os.Getenv("DB_PASSWORD")
	dbName := os.Getenv("DB_NAME")

	dbString := fmt.Sprintf("%s:%s@tcp(localhost:3306)/%s?charset=utf8&parseTime=True&loc=Local", user, password, dbName)

	conn, err := gorm.Open(mysql.Open(dbString), &gorm.Config{})
	if err != nil {
		return nil, err
	}
	return &DB{db: conn}, nil
}

// Returns the database instance
func (d *DB) GetDB() *gorm.DB {
	return d.db
}
