package main

import (
	"command_line_app/app"
	"fmt"
	"log"
	"os"
)

func main() {
	fmt.Println("Starting the program...")

	application := app.Genarate()
	if erro := application.Run(os.Args); erro != nil {
		log.Fatal(erro)
	}

}
