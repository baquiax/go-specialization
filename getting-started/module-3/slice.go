package main

import (
	"fmt"
	"sort"
	"strconv"
	"strings"
)

func main() {
	var input string

	slice := make([]int, 0, 3)

	fmt.Print("Please enter your numbers to order. Insert X to close the program.\n\n")

	for {
		fmt.Scan(&input)
		if strings.ToLower(input) == "x" {
			break
		}
		number, error := strconv.Atoi(input)

		if error != nil {
			fmt.Print("Please, enter a valid number\n\n")
			continue
		}

		slice = append(slice, number)
		sort.Ints(slice)

		fmt.Printf("Ordered slice: %v\n\n", slice)
	}
	fmt.Println("Good bye")
}
