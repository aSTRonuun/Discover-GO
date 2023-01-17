package main

import "fmt"

func avarage(numbers ...float64) float64 {
	total := 0.0
	for _, num := range numbers {
		total += num
	}
	return total / float64(len(numbers))
}

func main() {
	fmt.Printf("Avarage: %.2f", avarage(9.7, 3.1, 8.5, 9.1, 9.2))
}
