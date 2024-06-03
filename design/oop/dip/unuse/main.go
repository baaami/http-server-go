package main

// 의존 역전 원칙 (Dependency Inversion Principle, DIP) 미사용 예제

// Switch(고수준 모듈)는 LightBulb(저수준 모듈)에 직접 의존
// LightBulb를 다른 것으로 교체하려면 Switch 코드를 수정

import "fmt"

// Low-level module
type LightBulb struct{}

func (lb LightBulb) TurnOn() {
	fmt.Println("LightBulb: Turned on")
}

func (lb LightBulb) TurnOff() {
	fmt.Println("LightBulb: Turned off")
}

type Fan struct{}

func (f Fan) TurnOn() {
	fmt.Println("Fan: Turned on")
}

func (f Fan) TurnOff() {
	fmt.Println("Fan: Turned off")
}

// High-level module
type Switch struct {
	bulb    LightBulb
	fanBulb Fan
}

func (s Switch) Operate(turnOn bool) {
	if turnOn {
		s.bulb.TurnOn()
	} else {
		s.bulb.TurnOff()
	}
}

func (s Switch) OperateFan(turnOn bool) {
	if turnOn {
		s.fanBulb.TurnOn()
	} else {
		s.fanBulb.TurnOff()
	}
}

func main() {
	bulb := LightBulb{}
	fanBulb := Fan{}

	mySwitch := Switch{bulb: bulb}
	myFanSwitch := Switch{fanBulb: fanBulb}

	mySwitch.Operate(true)  // Turn on the light
	mySwitch.Operate(false) // Turn off the light

	myFanSwitch.OperateFan(true)
	myFanSwitch.OperateFan(false)
}
