package main

import "fmt"

func main() {

	channel := make(chan int, 2)

	channel <- 1 // send data to the channel (writing)
	channel <- 2

	<-channel // receving data to the channel (reading)
	fmt.Println(<-channel)
}
