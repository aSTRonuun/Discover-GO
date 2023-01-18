package main

import "fmt"

type sporty interface {
	turnOnTurbo()
}

type ferrari struct {
	model          string
	turboOn        bool
	actualVelocity int
}

func (f *ferrari) turnOnTurbo() {
	f.turboOn = true
}

func main() {
	car1 := ferrari{"F40", false, 0}
	car1.turnOnTurbo()

	var car2 sporty = &ferrari{"F80", false, 0}
	car2.turnOnTurbo()

	fmt.Println(car1, car2)
}
