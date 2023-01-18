package main

import "fmt"

type sporty interface {
	turnOnTurbo()
}

type luxurious interface {
	makeGoal()
}

type sportyLuxurious interface {
	sporty
	luxurious
}

type bwm7 struct{}

func (b bwm7) turnOnTurbo() {
	fmt.Println("Turbo...")
}

func (b bwm7) makeGoal() {
	fmt.Println("Make Goal...")
}

func main() {
	var b sportyLuxurious = bwm7{}
	b.makeGoal()
	b.turnOnTurbo()
}
