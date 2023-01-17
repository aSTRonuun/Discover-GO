package main

import "fmt"

func main() {
	s := make([]int, 10) // With a function make, we can make a slice and automatically the GO introces a reference between slice and array
	s[9] = 12
	fmt.Println(s)

	s = make([]int, 10, 20) // With a function make too, we can do the slice and say the array lenght
	fmt.Println(s, len(s), cap(s))

	s = append(s, 1, 2, 3, 4, 5, 6, 7, 8, 9, 0) // With funcion append, we can add more elements at the end of the array
	fmt.Println(s, len(s), cap(s))

	s = append(s, 1) // When we try add more elements at the end of the array and the array is limited in total, the GO automatically add more capacity in the array
	fmt.Println(s, len(s), cap(s))
}
