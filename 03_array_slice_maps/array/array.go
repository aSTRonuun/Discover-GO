package main

import "fmt"

func main() {
	// The array is homogeny (same type) and static (fixe)
	var grades [3]float64
	fmt.Println(grades)

	grades[0], grades[1], grades[2] = 7.2, 9.3, 8.9
	fmt.Println(grades)

	total := 0.0
	for i := 0; i < len(grades); i++ {
		total += grades[i]
	}

	avarage := total / float64(len(grades))
	fmt.Printf("Arage %.2f", avarage)
}
