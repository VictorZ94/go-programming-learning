package main

import (
	"fmt"
	"reflect"
)

func main() {
	curso := "Go's professional corse"
	fmt.Println(curso)

	// Get len of a string
	length := len(curso)
	fmt.Println(length)

	// to get position
	firstchar := curso[0]
	chars := curso[0:5]
	lastchar := curso[len(curso)-1]
	fmt.Println(firstchar)
	fmt.Println(chars)
	fmt.Println(lastchar)
	fmt.Println(reflect.TypeOf(firstchar))
	fmt.Printf("%c\n", firstchar)
	fmt.Printf("%c\n", lastchar)

}
