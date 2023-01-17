package main

import "fmt"

func main() {

	//var approveds map[int]string - Maps cannot just be only declared, they need to be declared and initialized
	approveds := make(map[int]string)

	approveds[2382382] = "Maria"
	approveds[2387239] = "Jhon"
	approveds[3327823] = "Cortana"

	fmt.Println(approveds)

	for cpf, name := range approveds {
		fmt.Printf("%s (CPF: %d)\n", name, cpf)
	}

	fmt.Println(approveds[3327823])
	delete(approveds, 3327823)

	fmt.Println(approveds)
}
