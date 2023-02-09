package routers
import (
	"github.com/gin-gonic/gin"
	"github.com/glitchdawg/week2-GL1-CipherSchools/database"
	"github.com/glitchdawg/week2-GL1-CipherSchools/handler"

)
func Router() *gin.Engine {
	router:=gin.Default()//get the default engine for further customization 
	api:= handler.Handler{
		DB: database.GetDB(),// set the handler DB
	}
	router.GET("/books", api.GetBooks)//set the http://localhost:8000
	//while calling handler function, gin will pass *gin.Context in the handler function
	return router
}
