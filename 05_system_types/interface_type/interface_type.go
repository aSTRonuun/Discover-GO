package main

import "fmt"

type curse struct {
	name string
}

func main() {
	var something interface{}
	fmt.Println(something)

	something = 3
	fmt.Println(something)

	type dynamic interface{} // We can use the interface type as a generic type
	var something2 dynamic = "Wow"
	fmt.Println(something2)

	something2 = true
	fmt.Println(something2)

	something2 = curse{"Golang: Exploring the Google language"}
	fmt.Println(something2)
}
