package main

import (
	"fmt"
	"os"
)

func main() {

	// if-else with inline declaration
	if i, err := fmt.Printf("random string\n"); err != nil {
		fmt.Println("Error occurred")
		os.Exit(1)
	} else if i > 100 {
		fmt.Printf("Bytes more than expected: %d", i)
	} else {
		fmt.Printf("Bytes returned: %d\n", i)
	}

	// simple error check
	_, err := fmt.Printf("another string\n")
	if err != nil {
		os.Exit(-1)
	}

}
