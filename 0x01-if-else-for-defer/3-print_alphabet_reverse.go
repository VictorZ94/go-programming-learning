package main

import "fmt"

func main() {
	a := 97 // 97 is same to "a" in ascii code
	defer fmt.Printf("\n")
	for a <= 122 {
		defer fmt.Printf("%c", byte(a))
		a++
	}
}
