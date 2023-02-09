package main

import (
	"log"

	"github.com/glitchdawg/week2-GL1-CipherSchools/database"
	"github.com/glitchdawg/week2-GL1-CipherSchools/routers"
)

func main() {
	database.Setup()
	engine:=routers.Router()
	err:=engine.Run("127.0.0.1:8000")
	if err!=nil{
		log.Fatal(err)
	}
}
