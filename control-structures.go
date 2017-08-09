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

	// switch: count vowels and consonants
	atoz := "the quick brown fox jumps over the lazy dog"
	vowels, consonants, zeds := 0, 0, 0

	for _, r := range atoz {
		switch r {
		case 'a', 'e', 'i', 'o', 'u':
			vowels++
		case 'z':
			zeds++
			fallthrough
		default:
			consonants++

		}
	}
	fmt.Printf("Vowels:%d, Consonants: %d (Zeds: %d)", vowels, consonants, zeds)

	// for
	for i, j := 0, 1; i <= 10; i, j = i+1, j*2 {
		fmt.Println(i, j)
	}

	// while
	var stop bool

	i := 0
	for !stop {
		fmt.Print(i)
		i++
		if i == 10 {
			stop = true
		}
	}

}
