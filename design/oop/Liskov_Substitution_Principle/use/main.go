package main

import "fmt"

type Bird interface {
	Walk()
}

type FlyingBird interface {
	Bird
	Fly()
}

type Sparrow struct{}

func (s Sparrow) Walk() {
	fmt.Println("Sparrow is walking")
}

func (s Sparrow) Fly() {
	fmt.Println("Sparrow is flying")
}

type Ostrich struct{}

func (o Ostrich) Walk() {
	fmt.Println("Ostrich is walking")
}

func main() {
	var bird Bird

	bird = Sparrow{}
	bird.Walk()

	var flyingBird FlyingBird
	flyingBird = Sparrow{}
	flyingBird.Fly()

	bird = Ostrich{}
	bird.Walk()
}
