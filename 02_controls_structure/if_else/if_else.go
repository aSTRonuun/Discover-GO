package main

import "fmt"

func printResult(grade float64) {
	if grade >= 7 {
		fmt.Println("Approved with grade", grade)
	} else {
		fmt.Println("Repproved with grade", grade)
	}
}

func main() {
	printResult(7.1)
	printResult(2.4)
}
