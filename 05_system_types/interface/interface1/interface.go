package main

import "fmt"

type printible interface {
	toString() string
}

type person struct {
	name    string
	surName string
}

type product struct {
	name  string
	price float64
}

type car struct {
	model string
	price float64
}

// The interfaces are impliciment implemented
func (p person) toString() string {
	return p.name + " " + p.surName
}

func (p product) toString() string {
	return fmt.Sprintf("%s - R$ %.2f", p.name, p.price)
}

func print(x printible) {
	fmt.Println(x.toString())
}

func main() {
	var something printible = person{"Roberto", "Claudio"}
	fmt.Println(something.toString())
	print(something)

	product := product{"Jeans", 99.99}
	print(product)

	// car := car{"Tesla", 29.999}
	// print(car) car cannot use the print function because, it doesn't implement the  toString() function
}
