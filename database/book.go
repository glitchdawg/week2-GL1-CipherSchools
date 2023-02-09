package database

import (
	"github.com/glitchdawg/week2-GL1-CipherSchools/models"
	"gorm.io/gorm"
)

func GetBooks(db *gorm.DB) ([]models.Book, error) {
	books := []models.Book{}
	query := db.Select("books.*")
	err := query.Find(&books).Error
	if err != nil {
		return nil, err
	}
	return books, nil
}
