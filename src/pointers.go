package main

import "fmt"

func main() {
	count := 10
	fmt.Printf("Value:%v, address: %v", count, &count)
}
