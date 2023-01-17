package main

import "fmt"

func main() {
	employeeByLetter := map[string]map[string]float64{
		"G": {
			"Gabi Nascimento": 27832.28,
			"Guda Reis":       382938.88,
		},
		"J": {
			"Joao Felix":   38298.22,
			"Jota Pereira": 82983.11,
		},
		"P": {
			"Pedro Fidelix": 30222.11,
			"Patati Arroda": 382.232,
		},
	}

	delete(employeeByLetter, "P")

	for letter, employeers := range employeeByLetter {
		for name, salary := range employeers {
			fmt.Printf("%s - %s Salary: %v \n", letter, name, salary)
		}
	}
}
