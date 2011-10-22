package main

import (
	"fmt"
	"sync"
)

type Counter struct {
	counter int
	sync.Mutex
}

func (self *Counter) Increment() {
	self.Lock()
	self.counter++
	self.Unlock()
}

func main() {
	printed := make(chan int)
	c := new(Counter)
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
