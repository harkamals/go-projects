package self_learn

import (
	"fmt"
	"reflect"
)

func init() {
	fmt.Println("Init called")
}

func main() {

	// fmt.Println(CompareNumbers(1, 2))

	// ARRAY
	arr := [5]int{1, 2, 3, 4, 5}
	fmt.Println(arr)
	fmt.Println(arr[0])

	for i := range arr {
		fmt.Println(i)
	}

	fmt.Println(reflect.TypeOf(arr).Elem())

	// SLICE
	my_slice := []int{1, 2, 3, 4, 5, 6, 7, 8}
	fmt.Println(my_slice)
	fmt.Println(my_slice[0])

	fmt.Printf("Length: %d, Capacity: %d", len(my_slice), cap(my_slice))

	my_slice = append(my_slice, 6, 7, 8, 9)

	for i := range my_slice {
		fmt.Println(i)
	}

	fmt.Printf("Length: %d, Capacity: %d", len(my_slice), cap(my_slice))
}

func CompareNumbers(a, b int) (bool, int) {

	/*
		if a > b {
			fmt.Println("a > b")
			return false, a - b
		} else if b > a {
			fmt.Println("b > a")
			return true, b - a
		}

		fmt.Println("a=b")
		return true, 0
	*/

	switch {
	case a > b:
		fmt.Println("a > b")
		return false, a - b
	case b > a:
		fmt.Println("b > a")
		return true, b - a

	}
	fmt.Println("a=b")
	return true, 0

}
