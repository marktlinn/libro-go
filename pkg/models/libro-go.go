package models

import (
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
	db := config.Connect().GetDB()
	db.AutoMigrate(&Book{})
}
