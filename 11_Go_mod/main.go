package main

import (
	"fmt"
	"module/auxiliar"

	"github.com/badoux/checkmail"
)

func main() {
	fmt.Println("Writing file from main")
	auxiliar.Writing()

	erro := checkmail.ValidateFormat("dev@gmail.com")
	fmt.Println(erro)
}
