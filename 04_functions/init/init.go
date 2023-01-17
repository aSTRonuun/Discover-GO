package main

import "fmt"

// The init function is a GO convention to initialize more first than main function
func init() {
	fmt.Println("Initializing...")
}

func main() {
	fmt.Println("Main...")
}
