package main

import (
	"fmt"
	"strings"
)

type person struct {
	name    string
	surName string
}

func (p person) getCompleteName() string {
	return p.name + " " + p.surName
}

func (p *person) setCompleteName(completeName string) {
	slices := strings.Split(completeName, " ")
	p.name = slices[0]
	p.surName = slices[1]
}

func main() {
	p1 := person{"Pedro", "Alves"}
	fmt.Println(p1.getCompleteName())

	p1.setCompleteName("Maria Pereira")
	fmt.Println(p1.getCompleteName())
}
