package main

import "fmt"

func main() {
	employeeAndSalaries := map[string]float64{
		"Jose Joao":    8388.83,
		"Gabi Matos":   7372.88,
		"Biatriz Dias": 100038.87,
	}

	employeeAndSalaries["Dias Reis"] = 38382.23
	delete(employeeAndSalaries, "Rafael Galdino")

	for name, salary := range employeeAndSalaries {
		fmt.Println(name, salary)
	}
}
