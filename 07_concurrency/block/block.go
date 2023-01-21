package main

import (
	"fmt"
	"time"
)

func routine(c chan int) {
	time.Sleep(time.Second)
	c <- 1 // blocking operation
	fmt.Println("Only Afeter data was readed")
}

func main() {
	c := make(chan int) // unbuffered channel

	go routine(c)

	fmt.Println(<-c) // bloking operation
	fmt.Println("Was readed")
	fmt.Println(<-c) // deadlook
	fmt.Println("End")
}
