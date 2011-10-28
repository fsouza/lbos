package main

import "fmt"
import "sync"

// Really bad example with a Getter and a Setter
// It smells Java, I know. I promise to avoid this
// and you should promise it also ;)
type Person struct {
	name string
	sync.RWMutex
}

func (p *Person) SetName(name string) {
	p.Lock()
	p.name = name
	p.Unlock()
}

func (p *Person) GetName() string {
	p.RLock()
	name := p.name
	p.RUnlock()
	return name
}

func main() {
	p := &Person{name: "Francisco"}
	names := []string{"Mary", "John", "Peter", "Michael", "Thomas", "Robert", "Andrew", "Alan"}
	finish := make(chan int)

	for _, name := range names {
		go func() {
			p.SetName(name)
		}()

		go func() {
			fmt.Println(p.GetName())
		}()
	}

	go func() {
		for {
			if p.GetName() == names[len(names) - 1] {
				finish <- 1
			}
		}
	}()

	<-finish
}
