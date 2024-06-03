package main

import "fmt"

// Large interface
type Worker interface {
	Work()
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

type Robot struct{}

func (r Robot) Work() {
	fmt.Println("Robot is working")
}

func (r Robot) Eat() {
	fmt.Println("Robot can not eat")
}

func main() {
	var worker Worker = Human{}
	worker.Work()
	worker.Eat()

	var robot Worker = Robot{}
	robot.Work()
}
