package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

// Animal interface of animals
type Animal interface {
	Eat()
	Move()
	Speak()
}

// Cow type
type Cow struct{}

// Snake type
type Snake struct{}

// Bird type
type Bird struct{}

// Eat satisfying Animal interface
func (cow *Cow) Eat() {
	fmt.Println("grass")
}

// Eat satisfying Animal interface
func (cow *Bird) Eat() {
	fmt.Println("worms")
}

// Eat satisfying Animal interface
func (cow *Snake) Eat() {
	fmt.Println("mice")
}

// Move satisfying Animal interface
func (cow *Cow) Move() {
	fmt.Println("walk")
}

// Move satisfying Animal interface
func (cow *Bird) Move() {
	fmt.Println("fly")
}

// Move satisfying Animal interface
func (cow *Snake) Move() {
	fmt.Println("slither")
}

// Speak satisfying Animal interface
func (cow *Cow) Speak() {
	fmt.Println("moo")
}

// Speak satisfying Animal interface
func (cow *Bird) Speak() {
	fmt.Println("peep")
}

// Speak satisfying Animal interface
func (cow *Snake) Speak() {
	fmt.Println("hsss")
}

func newAnimal(animals *map[string]Animal, name, animalType string) {
	switch animalType {
	case "cow":
		(*animals)[name] = &Cow{}
		fmt.Println("Created it!")
	case "bird":
		(*animals)[name] = &Bird{}
		fmt.Println("Created it!")
	case "snake":
		(*animals)[name] = &Snake{}
		fmt.Println("Created it!")
	default:
		fmt.Println("Ups! Unknown animal tpye. Try again.")
	}
}

func query(animals *map[string]Animal, name, action string) {
	animal := (*animals)[name]

	if animal == nil {
		fmt.Println("I don't exist... yet")
		return
	}

	switch action {
	case "eat":
		animal.Eat()
	case "move":
		animal.Move()
	case "speak":
		animal.Speak()
	default:
		fmt.Println("Invalid action. Try again.")
	}
}

func main() {
	fmt.Println("Hello, this is an simpel example to how use methods an structs in Go.")
	fmt.Println("You can use this commands:")
	fmt.Println("> newanimal <name> <cow|bird|snake>")
	fmt.Println("> query <name> <eat|move|speak>")
	fmt.Printf("\nPress CTRL + c to close\n")
	scanner := bufio.NewScanner(os.Stdin)
	animals := map[string]Animal{}

	for {
		fmt.Printf("\n> ")
		scanner.Scan()
		input := scanner.Text()
		commands := strings.Split(input, " ")

		if len(commands) != 3 {
			fmt.Println("Please insert a valid command")
			continue
		}

		command := commands[0]
		name := commands[1]
		actionOrType := commands[2]

		if command != "newanimal" && command != "query" {
			fmt.Printf("Supported commands: [newanimal, query]. Try again")
			continue
		}

		switch command {
		case "newanimal":
			newAnimal(&animals, name, actionOrType)
		default:
			query(&animals, name, actionOrType)
		}
	}
}
