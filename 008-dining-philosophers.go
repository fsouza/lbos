package main

import (
	"fmt"
	"sync"
	"time"
)

type Fork struct {
	sync.Mutex
}

type Table struct {
	philosophers chan Philosopher
	forks []*Fork
}

func NewTable(forks int) *Table {
	t := new(Table)
	t.philosophers = make(chan Philosopher, forks - 1)
	t.forks = make([]*Fork, forks)
	for i := 0; i < forks; i++ {
		t.forks[i] = new(Fork)
	}
	return t
}

func (t *Table) PushPhilosopher(p Philosopher) {
	p.table = t
	t.philosophers <- p
}

func (t *Table) PopPhilosopher() Philosopher {
	p := <-t.philosophers
	p.table = nil
	return p
}

func (t *Table) RightFork(philosopherIndex int) *Fork {
	f := t.forks[philosopherIndex]
	return f
}

func (t *Table) LeftFork(philosopherIndex int) *Fork {
	f := t.forks[(philosopherIndex + 1) % len(t.forks)]
	return f
}

type Philosopher struct {
	name string
	index int
	table *Table
	fed chan int
}

func (p Philosopher) Think() {
	fmt.Printf("%s is thinking...\n", p.name)
	time.Sleep(3e9)
	p.table.PushPhilosopher(p)
}

func (p Philosopher) Eat() {
	p.GetForks()
	fmt.Printf("%s is eating...\n", p.name)
	time.Sleep(3e9)
	p.PutForks()
	p.table.PopPhilosopher()
	p.fed <- 1
}

func (p Philosopher) GetForks() {
	rightFork := p.table.RightFork(p.index)
	rightFork.Lock()

	leftFork := p.table.LeftFork(p.index)
	leftFork.Lock()
}

func (p Philosopher) PutForks() {
	rightFork := p.table.RightFork(p.index)
	rightFork.Unlock()

	leftFork := p.table.LeftFork(p.index)
	leftFork.Unlock()
}

func main() {
	table := NewTable(5)
	philosophers := []Philosopher{
		Philosopher{"Thomas Nagel", 0, table, make(chan int)},
		Philosopher{"Elizabeth Anscombe", 1, table, make(chan int)},
		Philosopher{"Martin Heidegger", 2, table, make(chan int)},
		Philosopher{"Peter Lombard", 3, table, make(chan int)},
		Philosopher{"Gottfried Leibniz", 4, table, make(chan int)},
	}

	for {
		for _, p := range philosophers {
			go func(p Philosopher){
				p.Think()
				p.Eat()
			}(p)
		}

		for _, p := range philosophers {
			<-p.fed
			fmt.Printf("%s was fed.\n", p.name)
		}
	}

}
