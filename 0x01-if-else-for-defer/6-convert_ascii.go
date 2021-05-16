package main

import (
	"fmt"
)

func main() {

	// character to ASCII
	char := 'a' // rune, not string
	ascii := int(char)
	fmt.Println(string(char), ": ", ascii)

	// ASCII to character
	asciiNum := 65 // Uppercase A
	character := string(asciiNum)
	fmt.Println(asciiNum, ": ", character)
}
