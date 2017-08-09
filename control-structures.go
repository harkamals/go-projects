package main

import (
	"fmt"
	"os"
)

func main() {

	// inline declaration
	if i, err := fmt.Printf("random string\n"); err != nil {
		fmt.Println("Error occurred")
		os.Exit(1)
	} else {
		fmt.Printf("Bytes returned: %d", i)
	}

}
