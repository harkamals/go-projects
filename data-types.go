package main

import (
	"fmt"
	"math/rand"
	"strings"
)

func main() {
	fmt.Println("Hello World")

	name := "Jason Bourne"
	age := 35

	fmt.Println("Name:", name, "Age:", age)
	fmt.Printf("Name: %v Age: %d\n", name, age)

	// Tabular StdOut
	fmt.Printf("%15v %4v\n", name, age)
	fmt.Printf("%15v %4v\n", name, age+rand.Intn(age))

	// String
	var hero = strings.Contains(name, "Bourne")
	fmt.Printf("Hero: %t\n", hero)

	atoz := "the quick brown fox jumps over the lazy dog\n"
	fmt.Printf("%s\n", atoz[10:])
	fmt.Printf("%d\n", len(atoz))

	for i, char := range atoz[:10] {
		fmt.Printf("%d %c\n", i, char)
	}

}
