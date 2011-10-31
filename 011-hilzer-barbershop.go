package main

import (
	"fmt"
	"time"
)

type Barber int

type Customer int

func (b Barber) CutHair(barberChair, sofa, line, paying chan Customer) {
	select {
	case <-barberChair:
		fmt.Println("Cutting hair...")
		time.Sleep(2e9)
		paying <- 1
	case <-paying:
		fmt.Println("Receiving money from customer...")
		time.Sleep(5e8)

		select {
		case barberChair <- (<-sofa):
			fmt.Println("Customer moving from sofa to barberChair")
		case sofa <- (<-line):
			fmt.Println("Customer moving from line to sofa")
		default:
			fmt.Println("Barber sleeping...")
			time.Sleep(1e9)
		}

	default:
		fmt.Println("Barber sleeping...")
		time.Sleep(1e9)
	}
}

func main() {
	barbersCount := 3

	finish := make(chan int)
	barberChair := make(chan Customer, barbersCount)
	sofa := make(chan Customer, 4)
	line := make(chan Customer, 20)
	paying := make(chan Customer, 1)
	barber := Barber(0)

	for i := 0; i < barbersCount; i++ {
		go func(b Barber, barberChair, sofa, line, paying chan Customer){
			for {
				b.CutHair(barberChair, sofa, line, paying)
			}
		}(barber, barberChair, sofa, line, paying)
	}

	custCount := 30
	for i := 0; i < custCount; i++ {
		go func(barberChair, sofa, line chan Customer){
			select {
			case barberChair <- 1:
				fmt.Println("Customer going to get a haircut...")
			case sofa <- 1:
				fmt.Println("Customer sitting in the sofa...")
			case line <- 1:
				fmt.Println("Customer standing in the line...")
			default:
				fmt.Println("Customer giving up...")
			}
		}(barberChair, sofa, line)
	}

	<-finish
}
