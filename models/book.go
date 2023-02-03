package models

//the type of import that was done in the tutorial is no longer supported
//we use the following import statement instead
import "gorm.io/gorm"

type Book struct {
	gorm.Model
	Author    string
	Name      string
	pageCount int
}
