package main

import "fmt"

func factorial(n int) (int, error) {
	switch {
	case n < 0:
		return -1, fmt.Errorf("number is invalid: %d", n)
	case n == 0:
		return 1, nil
	default:
		anteriorFactorial, _ := factorial(n - 1)
		return n * anteriorFactorial, nil
	}
}

func main() {
	result, _ := factorial(5)
	fmt.Println(result)

	_, err := factorial(-5)
	if err != nil {
		fmt.Println(err)
	}
}
