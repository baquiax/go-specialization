package main

import (
	"fmt"
	"os"
)

type name struct {
	fname string
	lname string
}

func main() {
	var input string

	namesSlice := make([]name, 0)

	fmt.Println("Whats's the file name to read? (data.txt is included in the program)")
	fmt.Scan(&input)
	fileDescriptor, openError := os.Open(input)

	if openError != nil {
		fmt.Printf("Oh! An error has happened: %v", openError)
		return
	}

	firstNameSlice := make([]byte, 0)
	lastNameSlice := make([]byte, 0)
	splitReached := false

	for {
		char := make([]byte, 1)
		_, error := fileDescriptor.Read(char)

		if error != nil {
			if error.Error() == "EOF" {
				fmt.Printf("\nThe files was read completely\n\n")
				break
			}

			fmt.Printf("Oh! An error has happened reading: %v\n", error)
			continue
		}

		switch char[0] {
		case 10:
			namesSlice = append(namesSlice, name{string(firstNameSlice), string(lastNameSlice)})
			splitReached = false
			firstNameSlice = make([]byte, 0)
			lastNameSlice = make([]byte, 0)
			continue
		case 32:
			splitReached = true
			continue
		default:
			if splitReached {
				lastNameSlice = append(lastNameSlice, char[0])
			} else {
				firstNameSlice = append(firstNameSlice, char[0])
			}
		}
	}

	fileDescriptor.Close()

	fmt.Printf("Reading the slice:\n\n")
	for i, name := range namesSlice {
		fmt.Printf("%d. %s %s\n", i+1, name.fname, name.lname)
	}

	fmt.Printf("\nGood bye!\n")
}
