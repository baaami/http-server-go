package main

// 의존 역전 원칙 (Dependency Inversion Principle, DIP)예제

// Switch는 Switchable이라는 인터페이스에 의존

// Switch는 구체적인 LightBulb에 의존하지 않게 되어, 다른 종류의 Switchable 장치를 쉽게 교체
// LightBulb는 이 인터페이스를 구현

import "fmt"

// Abstraction
type Switchable interface {
	TurnOn()
	TurnOff()
}

// Low-level module
type LightBulb struct{}

func (lb LightBulb) TurnOn() {
	fmt.Println("LightBulb: Turned on")
}

func (lb LightBulb) TurnOff() {
	fmt.Println("LightBulb: Turned off")
}

// Low-level module
// Fan 추가, Switch 코드는 변경하지 않고도 다양한 장치를 사용할 수 있음. 이는 의존 역전 원칙을 잘 적용한 예시

type Fan struct{}

func (f Fan) TurnOn() {
	fmt.Println("Fan: Turned on")
}

func (f Fan) TurnOff() {
	fmt.Println("Fan: Turned off")
}

// High-level module
type Switch struct {
	device Switchable
}

func (s Switch) Operate(turnOn bool) {
	if turnOn {
		s.device.TurnOn()
	} else {
		s.device.TurnOff()
	}
}

func main() {
	bulb := LightBulb{}
	fan := Fan{}

	mySwitch := Switch{device: bulb}
	myFanSwitch := Switch{device: fan}

	mySwitch.Operate(true)  // Turn on the light
	mySwitch.Operate(false) // Turn off the light

	myFanSwitch.Operate(true)  // Turn on the fan
	myFanSwitch.Operate(false) // Turn off the fan
}
