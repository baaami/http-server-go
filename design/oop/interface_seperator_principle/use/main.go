package main

import "fmt"

// Small, specific interfaces
type Worker interface {
	Work()
}

type Eater interface {
	Eat()
}

// Worker implementation
type Human struct{}

func (h Human) Work() {
	fmt.Println("Human is working")
}

func (h Human) Eat() {
	fmt.Println("Human is eating")
}

// Another worker implementation that does not eat
type Robot struct{}

func (r Robot) Work() {
	fmt.Println("Robot is working")
}

func main() {
	var worker Worker = Human{}
	worker.Work()

	var eater Eater = Human{}
	eater.Eat()

	var robot Worker = Robot{}
	robot.Work()
}
