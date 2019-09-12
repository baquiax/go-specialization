package main

import (
	"fmt"
	"strconv"
	"strings"
)

func sort(orgArray []int, channel chan []int) {
	array := make([]int, len(orgArray))

	for i := 0; i < len(orgArray); i++ {
		array[i] = orgArray[i]
	}

	fmt.Printf("goroutine: array to sort %v\n", array)

	for i := 0; i < len(array); i++ {
		for j := 0; j < len(array)-1; j++ {
			if array[j] > array[j+1] {
				aux := array[j+1]
				array[j+1] = array[j]
				array[j] = aux
			}
		}
	}

	channel <- array
}

func main() {
	fmt.Println("Hi! Please insert one by one the numbers to be sorted (at least 4). Whe you'll finish, please write \"done\"")
	var input string
	numbers := make([]int, 0)

	for {
		fmt.Print("> ")
		fmt.Scan(&input)
		if strings.ToLower(input) == "done" {
			if len(numbers) < 4 {
				fmt.Printf("Please insert at least 4. %d more to go.\n", 4-len(numbers))
				continue
			}
			break
		}

		number, castError := strconv.Atoi(input)

		if castError != nil {
			fmt.Println("Please insert a valid number")
		}

		numbers = append(numbers, number)
	}

	channel := make(chan []int)
	totalNumbers := len(numbers)
	subArraysSize := totalNumbers / 4

	go sort(numbers[0:subArraysSize], channel)
	go sort(numbers[subArraysSize:subArraysSize*2], channel)
	go sort(numbers[subArraysSize*2:subArraysSize*3], channel)
	go sort(numbers[subArraysSize*3:], channel)

	first := <-channel
	second := <-channel
	third := <-channel
	fourth := <-channel

	partialSorted := append(first, second...)
	partialSorted = append(partialSorted, third...)
	partialSorted = append(partialSorted, fourth...)

	fmt.Printf("\nPartial sorted %v\n\n", partialSorted)

	go sort(partialSorted, channel)
	finalSort := <-channel

	fmt.Printf("Sorted %v\n\n", finalSort)
}
