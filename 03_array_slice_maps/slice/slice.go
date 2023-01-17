package main

import (
	"fmt"
	"reflect"
)

func main() {

	a1 := [3]int{2, 4, 6}
	s1 := []int{1, 3, 5}

	fmt.Println(a1, s1)
	fmt.Println(reflect.TypeOf(a1), reflect.TypeOf(s1))

	a2 := [5]int{1, 2, 3, 4, 5}

	// Slice isnt a array, Slice define a slice of a array
	s2 := a2[1:3]
	fmt.Println(a2, s2)

	// New slice, but is appoint to same array
	s3 := a2[:2]
	fmt.Println(a2, s3)

	// You can imagine a slice as: lenght and a point to a element of an array
	s4 := s2[:1]
	fmt.Println(s2, s4)
}
