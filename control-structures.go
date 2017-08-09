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
	n, err := fmt.Printf("switch case check..\n")

	switch {
	case err != nil:
		os.Exit(-1)
	case n == 0:
		fmt.Printf("no bytes recieved\n")
	case n != 18:
		fmt.Printf("n is not 10 bytes, it is: %d\n", n)
		fallthrough
	default:
		fmt.Println("Switch case: OK")
	}

	// switch: count vowels
	atoz := "the quick brown fox jumps over the lazy dog"
	vowels := 0

	for _, r := range atoz {
		switch r {
		case 'a', 'e', 'i', 'o', 'u':
			vowels++

		}

	}

	fmt.Printf("Vowels:%d", vowels)

}
