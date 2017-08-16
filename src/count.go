// MIT License
// Copyright (c) 2017 @harkamals

package src

import (
	"fmt"
	"time"
)

func printCount(c chan int) {

	num := 0
	for num >= 0 {
		num = <-c
		fmt.Println(num, " <-")
	}

}

func main() {

	fmt.Println("- start counter -")

	c := make(chan int)
	a := []int{8, 7, 6, 5, 4, 3, 2, 1, 0, 8, -1}

	go printCount(c)

	for _, v := range a {
		c <- v
	}

	time.Sleep(time.Millisecond)
	fmt.Println("- the end -")

}
