package main

import (
	"fmt"
)

func main() {

	a := 9
	b := 6

	calculate("+", a, b)
	calculate("-", a, b)
	calculate("*", a, b)
	calculate("/", a, b)

}

func calculate(op string, a, b int) (string, int) {
	result := 0

	switch op {
	case "+":
		result = a + b
	case "-":
		result = a - b
	case "*":
		result = a * b
	case "/":
		result = a / b
	}

	fmt.Printf("%d %v %d = %d\n", a, op, b, result)
	return op, result

}
