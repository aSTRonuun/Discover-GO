package main

import "fmt"

// Go Doesnt have tenary operator
func getResult(grade float64) string {
	if grade >= 7 {
		return "Approved"
	}
	return "Dissaproved"
}

func main() {
	fmt.Println(getResult(9.1))
}
