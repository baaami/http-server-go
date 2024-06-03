package main

import "fmt"

type Shape interface {
	Area() float64
}

type Rectangle struct {
	Width, Height float64
}

func (r Rectangle) Area() float64 {
	return r.Width * r.Height
}

type Circle struct {
	Radius float64
}

func (c Circle) Area() float64 {
	return 3.14 * c.Radius * c.Radius
}

// 삼각형 추가 시 다음과 같이 추가 가능
type Triangle struct {
	Base, Height float64
}

func (t Triangle) Area() float64 {
	return 0.5 * t.Base * t.Height
}

type AreaCalculator struct{}

func (ac AreaCalculator) CalculateTotalArea(shapes []Shape) float64 {
	var totalArea float64
	for _, shape := range shapes {
		totalArea += shape.Area()
	}
	return totalArea
}

func main() {
	shapes := []Shape{
		Rectangle{Width: 3, Height: 4},
		Circle{Radius: 5},
		Triangle{Base: 6, Height: 7},
	}
	calculator := AreaCalculator{}
	fmt.Println("Total Area:", calculator.CalculateTotalArea(shapes))
}
