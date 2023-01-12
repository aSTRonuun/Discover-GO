package main

import "fmt"

func main() {
	i := 1

	// Go doesnt have arithmetics of pointers
	var pointer *int = nil

	pointer = &i

	*pointer++
	i++
	fmt.Println(pointer, *pointer, i, &i)
}
