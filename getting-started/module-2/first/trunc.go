package main

import "fmt"

func main() {
	var floatNumber float32
	fmt.Println("Please enter a floating point number: ")
	fmt.Scan(&floatNumber)
	fmt.Printf("Truncated number: %d\n", int32(floatNumber))
}
