package main

import (
	"encoding/json"
	// "fmt"
	"io/ioutil"
	"log"
	"os"
)
type Something struct {
	Name    string  `json:"name"`
	Married bool    `json:"married"`
	Age     float64 `json:"age"`
	Address Address `json:"address"`
	Marks   []int   `json:"marks"`
}
type Address struct {
	Street int    `json:"street"`
	City   string `json:"city"`
}
func main() {
	jsonFile, err := os.Open("something.json")
	if err != nil {
		log.Fatal(err)
	}
	defer jsonFile.Close()
	jsonByteValues, err := ioutil.ReadAll(jsonFile)
	if err != nil {
		log.Fatal(err)
	}
	var something Something
	json.Unmarshal(jsonByteValues,&something)//converting json data to struct
	// fmt.Println(string(jsonByteValues))//converting json data to string
	log.Println(something)
	newJsonByteValues, err := json.Marshal(something)
	if err != nil {
		log.Fatal(err)
	}
	os.WriteFile("some.json",newJsonByteValues,0644)
	//consuming rest api in go code
	//using the rest api
	//returned value ? =json
	//now you must know how to read json data in string format
	//read json in struct format - marshalling, unmarshalling

}
