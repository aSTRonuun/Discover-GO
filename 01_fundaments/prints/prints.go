package main

import "fmt"

func main() {
	fmt.Print("Mesma")
	fmt.Print(" Linha")

	fmt.Println(" Nova")
	fmt.Print("Linha")

	x := 3.1415

	// fmt.Println("O valor de x é " + x) - Não pode converter e concatenar automaticamente

	xs := fmt.Sprint(x)
	fmt.Println("O valor de x é " + xs)
	fmt.Println("O valor de x é", x)

	fmt.Printf("O valor de x é %.2f", x)

	a := 1
	b := 1.99999
	c := false
	d := "opa"
	fmt.Printf("\n%d %f %.1f %t %s", a, b, b, c, d) // explict formt
	fmt.Printf("\n%v %v %v %v", a, b, c, d)
}
