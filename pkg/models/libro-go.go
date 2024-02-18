package models

import (
	"log"

	"github.com/marktlinn/libro-go/pkg/config"
	"gorm.io/gorm"
)

type Book struct {
	gorm.Model
	Name        string `json:"name"`
	Author      string `json:"author"`
	Publication string `json:"publication"`
}

func init() {
	db, err := config.Connect()
	if err != nil {
		log.Fatalf("error connecting to db: %s", err)
	}

	err = db.GetDB().AutoMigrate(&Book{})
	if err != nil {
		log.Fatalf("error auto-migrating Book model: %s", err)
	}
}

// Creates given Book in the database
func (b *Book) CreateBook(db *gorm.DB) *Book {
	db.Create(&b)
	return b
}

// Gets all Books from the database
func GetAllBooks(db *gorm.DB) ([]Book, error) {
	var books []Book
	if err := db.Find(&books).Error; err != nil {
		return nil, err
	}
	return books, nil
}

// Gets the specified Book according to its ID from the database.
func GetBookByID(ID int64, db *gorm.DB) *Book {
	var book Book
	db.Where("ID=?", ID).Find(&book)
	return &book
}

// Deletes the specified book from the database.
func DeleteBook(ID int64, db *gorm.DB) Book {
	var book Book
	db.Where("ID=?", ID).Delete(book)
	return book
}
