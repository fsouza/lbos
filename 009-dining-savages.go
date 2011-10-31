package main

import (
	"fmt"
	"time"
)

const M = 50

type Serving int

func SavagesThread(savages int, servings chan Serving) {
	for i := 0; i < savages; i++ {
		go func(){
			for {
				<-servings
				fmt.Println("Eating...")
				time.Sleep(1e9)
			}
		}()
	}
}

func CookThread(cooks int, servings chan Serving) {
	for i := 0; i < cooks; i++ {
		go func(){
			for {
				fmt.Println("Cooking...")
				time.Sleep(2e9)
				servings <- 1
			}
		}()
	}
}

func main() {
	finish := make(chan int)
	servings := make(chan Serving, M)
	CookThread(60, servings)
	SavagesThread(30, servings)

	<-finish
}
