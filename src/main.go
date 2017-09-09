package main

import (
	"fmt"
)

func main() {

	// Slice
	// slice_1 := make([]int, 5)
	slice_1 := []int{1, 2, 3, 4, 5}
	print_slice_stat(slice_1)

	slice_2 := append(slice_1, 6, 7, 8)
	print_slice_stat(slice_2)

	// Calculator
	a := 9
	b := 6

	calculate("+", a, b)
	calculate("-", a, b)
	calculate("*", a, b)
	calculate("/", a, b)

}
func print_slice_stat(slice []int) {
	fmt.Printf(
		"Capacity: %d, Len: %d, Value: %d\n",
		cap(slice),
		len(slice),
		slice)

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
