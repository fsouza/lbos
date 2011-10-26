package main

import (
	"fmt"
	"sync"
)

type Incrementer struct {
	counter int
	goal int
	waiting chan int
	waitingCount int
	waitingLimit int

	sync.Mutex
}

func NewIncrementer(barrierLimit, goal int) *Incrementer {
	i := new(Incrementer)

	if goal < barrierLimit {
		i.waitingLimit = goal
	} else {
		i.waitingLimit = barrierLimit
	}

	i.goal = goal
	i.waiting = make(chan int)
	return i
}

func (i *Incrementer) Wait() {
	i.Lock()

	if i.waitingLimit == i.waitingCount {
		for ; i.waitingCount > 0; i.waitingCount-- {
			i.waiting <- 1
		}
	}

	i.waitingCount++
	i.Unlock()

	fmt.Println("Waiting...")
	<-i.waiting
}

func (i *Incrementer) Signal(finish chan int) {
	if i.counter == i.goal {
		finish <- i.counter
	}
}

func (i *Incrementer) Increment(finish chan int) {
	i.Wait()
	fmt.Println("Incrementing...")
	i.counter++
	i.Signal(finish)
}

func (i *Incrementer) Run(finish chan int) {
	for j := 0; j < i.goal + 1; j++ {
		go i.Increment(finish)
	}
}

func main() {
	finish := make(chan int)
	i := NewIncrementer(10, 30)
	i.Run(finish)
	fmt.Println(<-finish)
}
