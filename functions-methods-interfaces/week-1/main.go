package main

import (
	"fmt"
	"strconv"
	"strings"
)

const maxNumbers = 10

// Swap swap the number int he position `index` with the
// number in the position `index + 1`
func Swap(numbersToSort []int, index int) {
	aux := numbersToSort[index]
	numbersToSort[index] = numbersToSort[index+1]
	numbersToSort[index+1] = aux
}

// BubbleSort One of the most simple algorithms of sorting
func BubbleSort(numbersToSort []int) {
	nElements := len(numbersToSort)

	for i := 1; i < nElements; i++ {
		for j := 0; j < nElements-1; j++ {
			if numbersToSort[j] > numbersToSort[j+1] {
				Swap(numbersToSort, j)
			}
		}
	}

}

func main() {
	var input string
	slicesOfNumbers := make([]int, 0)

	fmt.Printf("Hey, please insert %d numbers to sort:\n\n", maxNumbers)
	for {
		fmt.Scan(&input)

		if strings.ToLower(input) == "q" {
			fmt.Println("Bye")
			return
		}

		integer, numberError := strconv.Atoi(input)

		if numberError != nil {
			fmt.Printf("A number is required, please try again\n\n")
			continue
		}

		slicesOfNumbers = append(slicesOfNumbers, integer)

		if len(slicesOfNumbers) == maxNumbers {
			break
		}
	}

	BubbleSort(slicesOfNumbers)

	fmt.Printf("\nThe result is:\n")

	for i := 0; i < len(slicesOfNumbers); i++ {
		fmt.Printf("%d ", slicesOfNumbers[i])
	}

	fmt.Printf("\n\nBye\n")
}
