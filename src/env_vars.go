package main

import (
	"fmt"
	"os"
)

func main() {
	fmt.Println("hello")

	for i, e := range os.Environ() {
		fmt.Println(i, e)
	}

}
