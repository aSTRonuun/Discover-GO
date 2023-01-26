package main

import "fmt"

func recoverFromPanic() {
	if r := recover(); r != nil {
		fmt.Println("Recovering from panic")
	}
}

func IsAprroved(grades ...float64) bool {
	defer recoverFromPanic()

	total := 0.0
	for _, grade := range grades {
		total += grade
	}
	avarage := total / float64(len(grades))
	if avarage > 6 {
		return true
	} else if avarage < 6 {
		return false
	}

	// The panic function stops the program execution
	panic("avarage is 6")
}

func main() {
	fmt.Println(IsAprroved(6, 6))
}
