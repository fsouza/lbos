package main

import (
	"fmt"
	"sync"
	"time"
)

type List struct {
	insertMutex *sync.Mutex
	searchMutex *sync.RWMutex
	current int
	elements []int
}

func (s *List) Search(n int) (found bool) {
	s.searchMutex.RLock()
	for i := 0; i < len(s.elements) && !found; i++ {
		element := s.elements[i]
		found = (element == n)
	}
	s.searchMutex.RUnlock()

	return found
}

func (s *List) Insert(n int) {
	s.insertMutex.Lock()
	if len(s.elements) == cap(s.elements) && s.current < len(s.elements) {
		s.elements[s.current] = n
		s.current++
	} else if len(s.elements) < cap(s.elements) {
		oldLen := len(s.elements)
		s.elements = s.elements[:oldLen+1]
		s.elements[oldLen] = n
	} else {
		s.elements = append(s.elements, n)
	}
	s.insertMutex.Unlock()
}

func (s *List) Delete(n int) {
	s.insertMutex.Lock()
	s.searchMutex.Lock()

	index := -1
	for i := 0; i < len(s.elements) && index < 0; i++ {
		if s.elements[i] == n {
			index = i
		}
	}

	if index > -1 {
		for i := index; i < len(s.elements) - 1; i++ {
			s.elements[i] = s.elements[i+1]
		}

		s.elements = s.elements[:len(s.elements)-1]
	}


	s.searchMutex.Unlock()
	s.insertMutex.Unlock()
}

func (s *List) String() string {
	return fmt.Sprintf("%v", s.elements)
}

func NewList(initialLen int) *List {
	l := new(List)
	l.insertMutex = new(sync.Mutex)
	l.searchMutex = new(sync.RWMutex)
	l.elements = make([]int, initialLen)
	return l
}

func main() {
	finish := make(chan int)
	list := NewList(50)
	numbersToInsert := 200

	for i := 1; i < numbersToInsert+1; i++ {
		go func(n int) {
			list.Insert(n)
		}(i)
	}

	numbersToDelete := []int{4, 5, 10, 36, 190, 192, 200}
	for _, n := range numbersToDelete {
		go func(n int) {
			list.Delete(n)
		}(n)
	}

	numbersToSearch := []int{3, 3000, 200, 210, 5, 7, 9, 10, 51, 39233, 87, 66}
	for _, n := range numbersToSearch {
		go func(n int) {
			list.Search(n)
		}(n)
	}

	for {
		select {
		case <- finish:
		default:
			fmt.Println("Waiting...")
			fmt.Println(list)
			time.Sleep(2e9)
		}
	}
}
