package main

import (
	"fmt"
	"math"
	"reflect"
)

func main() {
	// integer numbers
	fmt.Println(1, 2, 1000)
	fmt.Println("Literal inteiro é", reflect.TypeOf(32000))

	// no signal (only positives)... uint8 uint16 uint32 uint64
	var b byte = 255 // byte -> uint8
	fmt.Println("The byte is", reflect.TypeOf(b))

	// with signal... int8 int16 int32 int64
	i1 := math.MaxInt64
	fmt.Println("The maximun value of int is", i1)
	fmt.Println("The type of i1 is", reflect.TypeOf(i1))

	var i2 rune = 'a' // rine represent a unicode table mapper
	fmt.Println("The rune is", reflect.TypeOf(i2))
	fmt.Println(i2)

	// reais numbers (float32, float64)
	var x float32 = 49.99 // For variable be float32, you need do expliciment
	fmt.Println("The type of x is", reflect.TypeOf(x))
	fmt.Println("The type of x is", reflect.TypeOf(4.4)) // float 64

	// boolean
	boo := true
	fmt.Println("The type of bool is", reflect.TypeOf(boo))
	fmt.Println(!boo)

	// string
	s1 := "Hello, my name is Vitor"
	fmt.Println(s1 + "!")
	fmt.Println("The lenght of string is", len(s1))

	// string with multiple lines
	s2 := `Olá
	my
	name
	is
	Vitor
	`
	fmt.Println(s2)
	fmt.Println("The lenght of string is", len(s2))

	// char is not valid -> he is a int
	char := 'a'
	fmt.Println("The type of char is", reflect.TypeOf(char))
	fmt.Println(char)
}
