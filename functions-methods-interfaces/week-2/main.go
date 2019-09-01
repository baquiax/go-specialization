package main

import (
	"fmt"
	"strconv"
)

// GenDisplaceFn returns a function that calculates distance
func GenDisplaceFn(
	acceleration float64,
	initialVelocity float64,
	initialDisplacement float64) func(time float64) float64 {
	return func(time float64) float64 {
		return (0.5 * acceleration * time * time) + (initialVelocity * time) + initialDisplacement
	}
}

func main() {
	var input string
	var acceleration, initialVelocity, initialDisplacement, time float64
	var convertionError error

	fmt.Printf("Hello, we'll calculate the distance S of some system. So please enter the required values:\n\n")

	for {
		fmt.Printf("Aceleration: ")
		fmt.Scan(&input)
		acceleration, convertionError = strconv.ParseFloat(input, 64)
		if convertionError != nil {
			fmt.Printf("Oh, please enter a valid number\n\n")
			continue
		}
		break
	}

	for {
		fmt.Printf("Initial velocity: ")
		fmt.Scan(&input)
		initialVelocity, convertionError = strconv.ParseFloat(input, 64)
		if convertionError != nil {
			fmt.Printf("Oh, please enter a valid number\n\n")
			continue
		}
		break
	}

	for {
		fmt.Printf("Initial displacement: ")
		fmt.Scan(&input)
		initialDisplacement, convertionError = strconv.ParseFloat(input, 64)
		if convertionError != nil {
			fmt.Printf("Oh, please enter a valid number\n\n")
			continue
		}
		break
	}

	for {
		fmt.Printf("Now, please give me the elapse time: ")
		fmt.Scan(&input)
		time, convertionError = strconv.ParseFloat(input, 64)
		if convertionError != nil {
			fmt.Printf("Oh, please enter a valid number\n\n")
			continue
		}
		break
	}

	sFunction := GenDisplaceFn(acceleration, initialVelocity, initialDisplacement)
	distance := sFunction(time)

	fmt.Printf("\nDistance: %f\n", distance)
}
