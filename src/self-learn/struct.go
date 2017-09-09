package self_learn

import (
	"fmt"
)

type bill struct {
	flag    bool
	counter int
	pi      float32
}

type lisa struct {
	flag    bool
	counter int
	pi      float32
}

func (b lisa) age() int {
	return b.counter
}

func (b *lisa) increment_age() {
	b.counter++
}

func main() {

	// define struct literal
	e1 := bill{
		flag:    true,
		counter: 5,
		pi:      3.1254,
	}

	e2 := lisa{false, 24, 3.1254}

	fmt.Println(e1, e2)
	// {true 5 3.1254} {false 24 3.1254}

	// e1 = bill(e2)
	fmt.Println(e1, e2)
	// {false 24 3.1254} {false 24 3.1254}

	// call receiver method
	e2.age()
	fmt.Println(e2.counter)

	e2.increment_age()
	fmt.Println(e2.counter)

	// Compare Types: bill and lisa
	// fmt.Println(bill{} == lisa{})
	// invalid operation: bill literal == lisa literal (mismatched types bill and lisa)

	// Need to convert lisa to type bill, then compare
	fmt.Println(bill{} == bill(lisa{}))

	// Anonymous struct, can be compared 1:1
	anonymous_struct := struct {
		flag    bool
		counter int
		pi      float32
	}{
		flag:    true,
		counter: 5,
		pi:      3.1254,
	}

	fmt.Println(e1 == anonymous_struct)
}
