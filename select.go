package main

import (
	"fmt"
	"time"
)

func main(){
	helloCh:=make(chan string,1)
	goodbyeCh:=make(chan string,1)
	quitCh:=make(chan bool)
	go receiveMessage(helloCh,goodbyeCh,quitCh)
	go sendMessage(helloCh,"Hello World")
	time.Sleep(time.Second)
	go sendMessage(helloCh, "Good Bye world")
	result:=<-quitCh
	fmt.Println("message from quitCh=",result)
}
func sendMessage(ch chan<- string, message string) {
	ch<-message
}
func receiveMessage(helloCh, goodbyeCh <- chan string, quitCh chan<-bool){
	for{
		select{
		case message:=<-helloCh:
			fmt.Println("Message from helloCh: ", message)
		case message:=<-goodbyeCh:
			fmt.Println("Message from goodbyeCh: ", message)
		case <- time.After(time.Second*2):
			fmt.Println("Nothing receiving in 2 seconds:   Exiting the receiveMessage goroutine")
			quitCh<-true
			break
		}
	}
}