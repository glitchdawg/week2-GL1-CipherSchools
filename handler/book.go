package handler

import (
	"log"
	"net/http"
	"github.com/gin-gonic/gin"
	"github.com/glitchdawg/week2-GL1-CipherSchools/database"
	"gorm.io/gorm"
)
type Handler struct{
	DB *gorm.DB
}
func (h *Handler)GetBooks(in *gin.Context){
	books,err:=database.GetBooks(h.DB)
	if err!=nil{
		log.Println(err)
		in.JSON(http.StatusInternalServerError,err)
	}
	in.JSON(http.StatusOK, books)
}