package main

import (
	"fmt"
	"time"
)

func rest(number *int, with int) {
	time.Sleep(500 * time.Microsecond)
	*number = (*number) - with
	fmt.Printf("Current value after rest %d: %d\n", with, *number)
}

func sum(number *int, with int) {
	time.Sleep(500 * time.Microsecond)
	*number = (*number) + with
	fmt.Printf("Current value after sum %d: %d\n", with, *number)
}
func main() {
	initialNumber := 100
	fmt.Printf("Initial number: %d\n", initialNumber)
	go rest(&initialNumber, 2)
	go sum(&initialNumber, 5)

	var exit string
	fmt.Scan(&exit)
}
