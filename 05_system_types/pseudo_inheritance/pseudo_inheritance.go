package main

import "fmt"

type car struct {
	name           string
	actualVelocity int
}

type ferrari struct {
	car    // anonymouys form
	turbOn bool
}

func main() {
	f := ferrari{}
	f.name = "F80"
	f.actualVelocity = 300
	f.turbOn = true

	fmt.Printf("Is Ferrari %s with turb on? %v\n", f.name, f.turbOn)
	fmt.Println(f)
}
