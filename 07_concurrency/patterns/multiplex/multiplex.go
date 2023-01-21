package main

import (
	"fmt"

	"github.com/aSTRonuun/html"
)

func forward(origin <-chan string, destination chan string) {
	for {
		destination <- <-origin
	}
}

// multiplex - mix messages int the channel
func join(input1, input2 <-chan string) <-chan string {
	c := make(chan string)
	go forward(input1, c)
	go forward(input2, c)
	return c
}

func main() {
	channel := join(
		html.Title("https://www.cod3r.com.br", "https://youtube.com"),
		html.Title("https://www.google.com", "https://gameplayscassi.com.br/"),
	)
	fmt.Println(<-channel, "|", <-channel)
	fmt.Println(<-channel, "|", <-channel)
}
