package main

import (
	"fmt"
	"strings"
)

func main() {
	var input string
	fmt.Print("Please, enter a string: ")
	fmt.Scan(&input)
	input = strings.ToLower(input)
	var hasPrefix = strings.HasPrefix(input, "i")
	var hasSuffix = strings.HasSuffix(input, "n")
	var containsA = strings.Contains(input, "a")
	if hasPrefix && hasSuffix && containsA {
		fmt.Println("Found!")
	} else {
		fmt.Println("Not found!")
	}
}
