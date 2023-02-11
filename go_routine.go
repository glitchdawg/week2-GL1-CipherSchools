package main

import (
	"fmt"
	"sync"
	// "time"
)
func main() {
	var wait sync.WaitGroup
	counter:=20000
	wait.Add(counter)
	for i:=0;i<counter;i++{
		go hello(&wait, i)
	}
	
	// go hello(wait) //not correct
	defer wait.Wait()
	// time.Sleep(time.Second * 5)
}	
func hello(wait *sync.WaitGroup, counter int) {
	fmt.Println(counter)
	// go func(){
	// 	fmt.Println(10)
	// 	fmt.Println(19)
	// 	fmt.Println(13)
	// }()
	wait.Done()
}
