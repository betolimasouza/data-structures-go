package main

import "fmt"

type IProcess interface {
	process()
}

type Adapter struct {
	adaptee Adaptee
}

func (adaptee Adaptee) convert() {
	fmt.Println("Adaptee Convert Method")
}

type Adaptee struct {
	adapterType int
}

func (adapter Adapter) process() {
	fmt.Println("Adapter Process")
	adapter.adaptee.convert()
}

func main() {
	var processor IProcess = Adapter{}
	processor.process()
}
