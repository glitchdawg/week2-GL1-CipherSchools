package main

import (
	"fmt"
	"log"

	"github.com/glitchdawg/week2-GL1-CipherSchools/database"
	"github.com/glitchdawg/week2-GL1-CipherSchools/routers"
)
func init(){
	database.Setup()//establish the database connection
}
func init(){
	fmt.Println(1)
}
func init(){
	fmt.Println(2)
}
func main() {
	
	engine := routers.Router()
	err := engine.Run("127.0.0.1:8000")
	if err != nil {
		log.Fatal(err)
	}
}
