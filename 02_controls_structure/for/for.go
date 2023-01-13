package main

import (
	"fmt"
	"time"
)

func main() {

	// Go dosent have while

	i := 1
	for i <= 10 {
		fmt.Println(i)
		i++
	}

	for i := 0; i <= 20; i += 2 {
		fmt.Printf("%d ", i)
	}

	// Infinit Lace
	for {
		fmt.Println("Awals...")
		time.Sleep(time.Second)
	}
}
