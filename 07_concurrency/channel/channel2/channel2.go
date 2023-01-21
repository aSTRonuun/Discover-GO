package main

import (
	"fmt"
	"time"
)

// Channel is the way of communication between goroutines
// Channel is a type

func twoThreeFourTimes(base int, channel chan int) {
	time.Sleep(time.Second)
	channel <- 2 * base

	time.Sleep(time.Second)
	channel <- 3 * base

	time.Sleep(3 * time.Second)
	channel <- 4 * base
}

func main() {
	c := make(chan int)

	go twoThreeFourTimes(2, c)
	fmt.Println("A")

	a, b := <-c, <-c
	fmt.Println("B")
	fmt.Println(a, b)

	fmt.Println(<-c)
}
