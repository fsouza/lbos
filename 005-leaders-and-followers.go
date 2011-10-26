package main

import (
	"fmt"
)

func main() {
	lQueue := make(chan int, 10)
	fQueue := make(chan int, 1)
	finish := make(chan int)

	go func(){
		fmt.Println("Follower arrived")
		fQueue <- 1
		<-lQueue
		fmt.Println("Follower and leader leaving")
		finish <- 1
	}()

	go func(){
		fmt.Println("Leader arrived")
		lQueue <- 1
		<-fQueue
		fmt.Println("Leader and follower leaving")
		finish <- 1
	}()

	<-finish
}
