package main

import (
	"fmt"
	"strconv"
)

func main() {
	x := 2.4
	y := 2
	fmt.Println(x / float64(y))

	grade := 6.9

	finalGrade := int(grade)
	fmt.Println(finalGrade)

	// convert int to char
	fmt.Println("Test " + string(97))

	// convert int to string
	fmt.Println("Test " + strconv.Itoa(123))

	// string to int
	num, _ := strconv.Atoi("123")
	fmt.Println(num - 122)

	b, _ := strconv.ParseBool("true")
	if b {
		fmt.Println("True")
	}
}
