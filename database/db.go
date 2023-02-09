package database

import (
	"log"

	"github.com/glitchdawg/week2-GL1-CipherSchools/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)
var DB *gorm.DB
func GetDB() *gorm.DB{
	return DB
}
func Setup() {
	// host:="localhost"
	// port:="5432"
	// dbName:="book"
	// username:="postgres"
	// password:="postgres"
	// arg:="host="+host+" port="+port+" user="+username+" dbname"+dbName+" sslmode=disable password="+password
	// db,err:= gorm.Open("postgres",arg)
	dsn := "host=localhost user=postgres password=postgres dbname=book  port=9920 sslmode=disable TimeZone=Asia/Shanghai"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal(err)
	}
	db.AutoMigrate(&models.Book{})
	DB=db
}
 