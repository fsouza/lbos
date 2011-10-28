package main

import (
	"fmt"
)

type Executer interface {
	Execute()
}

type Buffer struct {
	Commands chan Executer
}

func NewBuffer() *Buffer {
	b := new(Buffer)
	b.Commands = make(chan Executer)
	return b
}

type NamePrinter struct {
	name string
}

func (n *NamePrinter) Execute() {
	fmt.Println(n.name)
}

func main() {
	buffer := NewBuffer()

	go func() {
		for {
			command := <-buffer.Commands
			command.Execute()
		}
	}()

	printer := &NamePrinter{"Francisco Souza"}
	anotherPrinter := &NamePrinter{"Anonymous"}

	buffer.Commands <- printer
	buffer.Commands <- anotherPrinter
}
