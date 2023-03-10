package main

import (
	"fmt"
	m "math"
)

func main() {
	const PI float64 = 3.1415

	var raio = 3.2

	// shorthand way to initialize variables
	area := PI * m.Pow(raio, 2)

	fmt.Println("A área da circuferência é", area)

	// initializing using blocks
	const (
		a = 1
		b = 2
	)

	var (
		c = 3
		d = 4
	)
	fmt.Println(a, b, c, d)

	// Initializing using a only line
	var e, f bool = true, false
	fmt.Println(e, f)

	g, h, i := 2, false, "epa!"
	fmt.Println(g, h, i)
}
