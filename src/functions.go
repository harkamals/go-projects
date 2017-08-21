// MIT License
// Copyright (c) 2017 @harkamals

package main

import (
	"fmt"
	"os"
)

func printer(format string, messages ...string) (count int, e error) {
	defer fmt.Printf("Messages: %d\n", len(messages))

	_, e = fmt.Printf(format, messages)
	count = len(messages)

	return
}

func main() {

	_, err := printer(
		"%s\n",
		"Hello", " World!", "How is the weather")

	if err != nil {
		os.Exit(-1)
	}

}
