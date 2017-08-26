package main

import (
	"fmt"
)

type bill struct {
	flag    bool
	counter int16
	pi      float32
}

type lisa struct {
	flag    bool
	counter int16
	pi      float32
}

func main() {

	// Compare Types: bill and lisa
	// fmt.Println(bill{} == lisa{})
	// invalid operation: bill literal == lisa literal (mismatched types bill and lisa)

	// Need to convert lisa to type bill, then compare
	fmt.Println(bill{} == bill(lisa{}))

	e1 := bill{
		flag:    true,
		counter: 5,
		pi:      3.1254,
	}

	e2 := lisa{false, 24, 3.1254}
	fmt.Println(e1, e2)
	// {true 5 3.1254} {false 24 3.1254}

	e1 = bill(e2)
	fmt.Println(e1, e2)
	// {false 24 3.1254} {false 24 3.1254}

}
