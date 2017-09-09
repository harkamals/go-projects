// MIT License
// Copyright (c) 2017 @harkamals

package self_learn

import (
	"fmt"
	"io/ioutil"
	"math/rand"
	"net/http"
	_ "net/url"
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

	// String - Literal
	line := `"nothing" gets escaped \n`
	fmt.Printf("Literal string: %s\n", line)

	// String
	var hero = strings.Contains(name, "Bourne")
	fmt.Printf("\nHero: %t\n", hero)

	atoz := "the quick brown fox jumps over the lazy dog\n"
	fmt.Printf("%s\n", atoz[10:])
	fmt.Printf("%d\n", len(atoz))

	for i, char := range atoz[:10] {
		fmt.Printf("%d %c\n", i, char)
	}

	// Array
	words := [...]string{"the", "quick", "brown", "fox"}
	fmt.Println(words)

	// Array - fixed length
	words4 := [5]string{"the", "quick", "brown", "fox"}
	fmt.Println(words4)

}
