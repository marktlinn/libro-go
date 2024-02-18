package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/marktlinn/libro-go/pkg/models"
	"gorm.io/gorm"
)

var NewBook models.Book

func CreateBook() {
	// TODO
}

func GetBooks(w http.ResponseWriter, r *http.Request, db *gorm.DB) {
	newBooks, err := models.GetAllBooks(db)
	if err != nil {
		http.Error(w, fmt.Sprintf("failed to get books %s", err), http.StatusInternalServerError)
	}

	if err = json.NewEncoder(w).Encode(newBooks); err != nil {
		http.Error(w, fmt.Sprintf("failed to encode response: %s", err), http.StatusInternalServerError)
	}

	w.WriteHeader(http.StatusOK)
}

func GetBookById() {
	// TODO

}

func UpdateBook() {
	// TODO

}

func DeleteBook() {
	// TODO
}
