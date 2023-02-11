package database

import (
	"github.com/glitchdawg/week2-GL1-CipherSchools/models"
	"gorm.io/gorm"
)
//get books is creating conection and interacting from golang app to db server via db variable
func GetBooks(db *gorm.DB) ([]models.Book, error) {
	books := []models.Book{}
	query := db.Select("books.*")
	err := query.Find(&books).Error
	if err != nil {
		return nil, err
	}
	return books, nil
}
func GetBookByID(db *gorm.DB, id string) (*models.Book, error) {
	book:=models.Book{}
	query := db.Select("books.*").Group("books.id").Where("books.id= ?",id).First(&book).Error
	if err!=nil{
		return nil, err
	}



	return models.Book{}, nil
}
func DeleteBookByID(db *gorm.DB, id string)error{
	var book models.Book
	err:=db.Where("id=?", id).Delete(&book).Error
	if err!=nil{
		return nil, err
	}
	return nil
}
func UpdateBookByID(db *gorm.DB, *models.Book)error{
	err:= db.Save(book).Error
	if err!=nil{
		return err
	}
	return nil
}
