package main

import "fmt"

func main() {

	s1 := make([]int, 10, 20)

	s2 := append(s1, 1, 2, 3)

	fmt.Println(s1, s2)

	// both slices point to the same refence array
	s1[0] = 2
	fmt.Println(s1, s2)
}
