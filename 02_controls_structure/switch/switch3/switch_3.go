package main

import (
	"fmt"
	"time"
)

func types(i interface{}) string {
	switch i.(type) {
	case int:
		return "integer"
	case float32, float64:
		return "real"
	case string:
		return "string"
	case func():
		return "function"
	default:
		return "Dont know"
	}
}

func main() {
	fmt.Println(types(2.1))
	fmt.Println(types(1))
	fmt.Println(types("Wow"))
	fmt.Println(types(func() {}))
	fmt.Println(types(time.Now()))
}
