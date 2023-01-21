package main

import (
	"fmt"
	"time"
)

func speak(person, text string, qtd int) {
	for i := 0; i < qtd; i++ {
		time.Sleep(time.Second)
		fmt.Printf("%s: %s (iteration %d)\n", person, text, i)
	}
}

func main() {
	//speak("Mary", "Why do not you talk to me", 5)
	//speak("Joao", "I only can speak after you", 1)

	// go speak("Mary", "Hey...", 500)
	// go speak("Joao", "Helo...", 500)

	go speak("Mary", "I understand!!!", 10)
	speak("Joao", "Congratulations", 5)
}
