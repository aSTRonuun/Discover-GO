package main

import "fmt"

type grade float64

func (g grade) between(begin, end float64) bool {
	return g >= grade(begin) && g <= grade(end)
}

func gradeToConcept(g grade) string {
	switch {
	case g.between(9.0, 10.0):
		return "A"
	case g.between(7.0, 8.99):
		return "B"
	case g.between(5.0, 6.99):
		return "C"
	case g.between(3.0, 4.99):
		return "D"
	default:
		return "E"
	}
}

func main() {
	fmt.Println(gradeToConcept(8.37))
	fmt.Println(gradeToConcept(9.37))
	fmt.Println(gradeToConcept(4.37))
}
