package main

import (
	"fmt"
	"time"
)

type Barber int

type Customer int

func (b Barber) CutHair(customers chan Customer) {
	select {
	case <-customers:
		fmt.Println("Cutting hair...")
		time.Sleep(2e9)
	default:
		fmt.Println("Sleeping...")
		time.Sleep(1e9)
	}
}

func main() {
	finish := make(chan int)
	customers := make(chan Customer, 5)
	barber := Barber(0)

	go func(b Barber, customers chan Customer){
		for {
			b.CutHair(customers)
		}
	}(barber, customers)

	custCount := 1200
	for i := 0; i < custCount; i++ {
		go func(customers chan Customer){
			select {
			case customers <- 1:
				fmt.Println("Customer going to get a haircut...")
			default:
				fmt.Println("Customer giving up...")
			}
		}(customers)
	}

	<-finish
}
