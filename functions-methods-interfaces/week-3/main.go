package main

import (
	"fmt"
	"strings"
)

// Animal defines a type with 3 values associated with an animal
type Animal struct {
	foodEaten        string
	locomotionMethod string
	spokenSound      string
}

// Eat prints the food eaten by the recevier animal
func (a *Animal) Eat() {
	fmt.Println(a.foodEaten)
}

// Move prints the locomotion method of the recevier animal
func (a *Animal) Move() {
	fmt.Println(a.locomotionMethod)
}

// Speak prints the spoken sound of the recevier animal
func (a *Animal) Speak() {
	fmt.Println(a.spokenSound)
}

// IsEmpty validate if the struct is a 0 value
func (a *Animal) IsEmpty() bool {
	return a.foodEaten == "" && a.locomotionMethod == "" && a.spokenSound == ""
}

func main() {
	cow := Animal{
		foodEaten:        "grass",
		locomotionMethod: "walk",
		spokenSound:      "moo",
	}

	bird := Animal{
		foodEaten:        "worms",
		locomotionMethod: "fly",
		spokenSound:      "peep",
	}

	snake := Animal{
		foodEaten:        "mice",
		locomotionMethod: "slither",
		spokenSound:      "hss",
	}

	var animal, action string

	knownAnimals := map[string]Animal{
		"cow":   cow,
		"bird":  bird,
		"snake": snake,
	}

	fmt.Println("Hello, this is an simpel example to how use methods an structs in Go.")
	fmt.Println("You can use this sintax <animal> <action>")
	fmt.Printf("Known animals: cow, bird and snake. And kwnown action: eat, speak and move\n\n")
	fmt.Printf("Press CTRL + c to close\n\n")
	for {
		fmt.Print("> ")
		_, scanError := fmt.Scanf("%s %s\n", &animal, &action)

		if scanError != nil {
			fmt.Printf("\nPlease enter a valid command like <animal> <action>\n\n")
			continue
		}

		chosenAnimal := knownAnimals[strings.ToLower(animal)]

		if chosenAnimal.IsEmpty() {
			fmt.Printf("\nI don't know nothing about %s, sorry try again\n\n", animal)
			continue
		}

		switch strings.ToLower(action) {
		case "eat":
			chosenAnimal.Eat()
		case "move":
			chosenAnimal.Move()
		case "speak":
			chosenAnimal.Speak()
		default:
			fmt.Printf("\nUps! The %s can't do it\n\n", animal)
		}
	}
}
