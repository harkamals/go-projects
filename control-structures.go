package main

import (
	"fmt"
	"os"
)

func main() {

	// simple error check
	_, err := fmt.Printf("random string\n")
	if err != nil {
		os.Exit(-1)
	}

	// if-else with inline declaration
	if i, err := fmt.Printf("another string\n"); err != nil {
		fmt.Println("Error occurred")
		os.Exit(1)
	} else if i > 100 {
		fmt.Printf("Bytes more than expected: %d", i)
	} else {
		fmt.Printf("Bytes returned: %d\n", i)
	}

	// switch
	n, err := fmt.Printf("switch case check\n")

	switch {
	case err != nil:
		os.Exit(-1)
	case n == 0:
		fmt.Printf("no bytes recieved\n")
	case n != 18:
		fmt.Printf("n is not 10 bytes, it is: %d\n", n)
	default:
		fmt.Printf("Switch case: OK")
	}

}
