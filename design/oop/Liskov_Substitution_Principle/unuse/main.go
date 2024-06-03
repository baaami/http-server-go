package main

import "fmt"

type Bird interface {
	Fly()
}

type Sparrow struct{}

func (s Sparrow) Fly() {
	fmt.Println("Sparrow is flying")
}

type Ostrich struct{}

func (o Ostrich) Fly() {
	fmt.Println("Ostrich can't fly")
}

func main() {
	var bird Bird

	bird = Sparrow{}
	bird.Fly()

	bird = Ostrich{}
	bird.Fly()
}
