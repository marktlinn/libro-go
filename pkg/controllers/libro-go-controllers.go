package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/marktlinn/libro-go/pkg/models"
	"github.com/marktlinn/libro-go/pkg/utils"
	"gorm.io/gorm"
)

const (
	FailedToEncodeRes = "failed to encode response"
	FailedToParseId   = "failed to parse book id"
)

func CreateBook(w http.ResponseWriter, r *http.Request, db *gorm.DB) {
	createdBook := &models.Book{}
	if err := utils.ParseBody(r, createdBook); err != nil {
		http.Error(w, fmt.Sprintf("failed to parse book data %s", err), http.StatusBadRequest)
		return
	}

	b, err := createdBook.CreateBook(db)

	if err != nil {
		http.Error(w, fmt.Sprintf("failed to create book %s", err), http.StatusInternalServerError)
		return
	}

	setHeaderOK(w)
	if err := json.NewEncoder(w).Encode(b); err != nil {
		http.Error(w, fmt.Sprintf("%s: %s", FailedToEncodeRes, err), http.StatusInternalServerError)
		return
	}
}

func GetBooks(w http.ResponseWriter, r *http.Request, db *gorm.DB) {
	newBooks, err := models.GetAllBooks(db)
	if err != nil {
		http.Error(w, fmt.Sprintf("failed to get books %s", err), http.StatusInternalServerError)
		return
	}

	setHeaderOK(w)
	if err = json.NewEncoder(w).Encode(newBooks); err != nil {
		http.Error(w, fmt.Sprintf("%s: %s", FailedToEncodeRes, err), http.StatusInternalServerError)
		return
	}
}

func GetBookById(w http.ResponseWriter, r *http.Request, db *gorm.DB) {
	vars := mux.Vars(r)
	idString := vars["id"]
	id, err := strconv.ParseInt(idString, 0, 0)
	if err != nil {
		http.Error(w, fmt.Sprintf("%s: %s", FailedToParseId, err), http.StatusBadRequest)
		return
	}

	book, err := models.GetBookByID(id, db)
	if err != nil {
		http.Error(w, fmt.Sprintf("failed to get book %s", err), http.StatusInternalServerError)
		return
	}

	setHeaderOK(w)
	if err = json.NewEncoder(w).Encode(book); err != nil {
		http.Error(w, fmt.Sprintf("%s: %s", FailedToEncodeRes, err), http.StatusInternalServerError)
		return
	}

}

func UpdateBook(w http.ResponseWriter, r *http.Request, db *gorm.DB) {
	updatedBook := &models.Book{}
	if err := utils.ParseBody(r, updatedBook); err != nil {
		http.Error(w, fmt.Sprintf("%s: %s", FailedToParseId, err), http.StatusBadRequest)
		return
	}
	vars := mux.Vars(r)
	idString := vars["id"]

	id, err := strconv.ParseInt(idString, 0, 0)
	if err != nil {
		http.Error(w, fmt.Sprintf("%s: %s", FailedToParseId, err), http.StatusBadRequest)
		return
	}

	b, err := models.GetBookByID(id, db)
	if err != nil {
		http.Error(w, fmt.Sprintf("Failed to get book by id: %s", err), http.StatusInternalServerError)
		return
	}

	b.UpdateBookData(updatedBook)
	db.Save(&b)

	setHeaderOK(w)
	if err = json.NewEncoder(w).Encode(b); err != nil {
		http.Error(w, fmt.Sprintf("%s: %s", FailedToEncodeRes, err), http.StatusInternalServerError)
		return
	}
}

func DeleteBook(w http.ResponseWriter, r *http.Request, db *gorm.DB) {
	vars := mux.Vars(r)
	idString := vars["id"]
	id, err := strconv.ParseInt(idString, 0, 0)
	if err != nil {
		http.Error(w, fmt.Sprintf("failed to parse book id: %s", err), http.StatusBadRequest)
		return
	}

	b, err := models.DeleteBook(id, db)
	if err != nil {
		http.Error(w, fmt.Sprintf("failed to delete book: %s", err), http.StatusBadRequest)
		return
	}

	setHeaderOK(w)
	if err = json.NewEncoder(w).Encode(b); err != nil {
		http.Error(w, fmt.Sprintf("%s: %s", FailedToEncodeRes, err), http.StatusInternalServerError)
		return
	}
}

// Sets the http.ResponseWriter to "Content-Type" = "application/json", with an http Status Code of 200
func setHeaderOK(w http.ResponseWriter) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
}
