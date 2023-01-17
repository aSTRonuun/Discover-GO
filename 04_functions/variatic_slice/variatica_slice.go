package main

import "fmt"

func printApproveds(approveds ...string) {
	fmt.Println("List of Approveds")
	for i, approved := range approveds {
		fmt.Printf("%d) %s\n", i+1, approved)
	}
}

func main() {
	approveds := []string{"Mary", "Pedro", "Victor", "Guilherme"}
	printApproveds(approveds...)
}
