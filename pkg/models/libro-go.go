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
func (b *Book) CreateBook(db *gorm.DB) (*Book, error) {
	if err := db.Create(&b).Error; err != nil {
		return nil, err
	}
	return b, nil
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
func GetBookByID(ID int64, db *gorm.DB) (*Book, error) {
	var book Book
	if err := db.Where("ID=?", ID).Find(&book).Error; err != nil {
		return nil, err
	}
	return &book, nil
}

// Deletes the specified book from the database.
func DeleteBook(ID int64, db *gorm.DB) (Book, error) {
	var book Book
	if err := db.First(&book, ID).Error; err != nil {
		return Book{}, err
	}

	if err := db.Where("ID=?", ID).Delete(&book).Error; err != nil {
		return Book{}, err
	}

	return book, nil
}

// Updates this book's fields with the newData passed.
func (b *Book) UpdateBookData(newData *Book) *Book {
	if newData.Author != "" {
		b.Author = newData.Author
	}
	if newData.Name != "" {
		b.Name = newData.Name
	}
	if newData.Publication != "" {
		b.Publication = newData.Publication
	}
	return b
}
