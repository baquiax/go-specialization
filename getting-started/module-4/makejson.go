package main

import (
	"encoding/json"
	"fmt"
)

func main() {
	var input string
	var person = map[string]string{}

	fmt.Print("Hello!\n\n")

	fmt.Println("What's your name?")
	fmt.Scan(&input)
	person["name"] = input

	fmt.Println("\nWhat's your address?")
	fmt.Scan(&input)
	person["address"] = input

	personAsJSON, _ := json.Marshal(person)

	fmt.Printf("\nHey this is your JSON: %s\n", personAsJSON)
}
