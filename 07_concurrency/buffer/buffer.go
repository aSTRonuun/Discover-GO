package main

import (
	"fmt"
	"time"
)

func routine(channel chan int) {
	channel <- 1
	channel <- 2
	channel <- 3
	channel <- 4
	channel <- 5
	fmt.Println("Executed")
	channel <- 6
}

func main() {

	ch := make(chan int, 3)

	go routine(ch)

	time.Sleep(time.Second)
	fmt.Println(<-ch)
}
