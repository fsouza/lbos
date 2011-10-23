package main

import (
	"fmt"
)

func ThreadA(aArrived chan<- int, bArrived <-chan int, finished chan<- int) {
	fmt.Println("a1")
	aArrived <- 1
	<-bArrived
	fmt.Println("a2")
	finished <- 1
}

func ThreadB(aArrived <-chan int, bArrived chan<- int, finished chan<- int) {
	fmt.Println("b1")
	bArrived <- 1
	<-aArrived
	fmt.Println("b2")
	finished <- 1
}

func main() {
	aArrived := make(chan int, 1)
	bArrived := make(chan int)
	finished := make(chan int)

	go ThreadA(aArrived, bArrived, finished)
	go ThreadB(aArrived, bArrived, finished)

	<-finished
}
