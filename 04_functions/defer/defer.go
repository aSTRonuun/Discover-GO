package main

import "fmt"

func getApprovedValue(number int) int {
	defer fmt.Println("Fim!") // Defer is used to excute a line of code before a function return

	if number > 5000 {
		fmt.Println("Big Value...")
		return number
	}
	fmt.Println("Small Value...")
	return number
}

func main() {
	fmt.Println(getApprovedValue(6000))
	fmt.Println(getApprovedValue(3000))
}
