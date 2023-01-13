package main

import (
	"fmt"
	"time"
)

func main() {

	t := time.Now()
	switch {
	case t.Hour() < 12:
		fmt.Println("Good Morning")
	case t.Hour() < 18:
		fmt.Println("Good Afeternoon")
	default:
		fmt.Println("Good Night")
	}
}
