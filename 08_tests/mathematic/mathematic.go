package mathematic

import (
	"fmt"
	"strconv"
)

func Avarage(numbers ...float64) float64 {
	total := 0.0
	for _, number := range numbers {
		total += number
	}

	avarage := total / float64(len(numbers))
	arroundAvarage, _ := strconv.ParseFloat(fmt.Sprintf("%.2f", avarage), 64)
	return arroundAvarage
}
