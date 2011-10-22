package main

import (
	"fmt"
)

type Counter struct {
	counter int
	multiplex chan int
}

func (self *Counter) Lock() {
	self.multiplex <- 1
}

func (self *Counter) Unlock() {
	<-self.multiplex
}

func (self *Counter) Increment() {
	self.Lock()
	self.counter++
	self.Unlock()
}

func NewCounter(multiplexLimit int) *Counter {
	counter := new(Counter)
	counter.multiplex = make(chan int, multiplexLimit)
	return counter
}

func main() {
	printed := make(chan int)
	c := NewCounter(10)
	go c.Increment()
	go c.Increment()
	go c.Increment()
	go c.Increment()
	go c.Increment()
	go c.Increment()
	go c.Increment()
	go c.Increment()
	go c.Increment()
	go c.Increment()
	go c.Increment()
	go c.Increment()
	go c.Increment()
	go c.Increment()
	go c.Increment()
	go c.Increment()
	go func(c *Counter){
		c.Lock()
		fmt.Println(c.counter)
		c.Unlock()
		printed <- 1
	}(c)

	<-printed
}

