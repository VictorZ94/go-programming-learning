package main

import "fmt"

func main() {
	a := 97 // 97 is same to "a" in ascii code
	for a <= 122 {
		fmt.Printf("%c", byte(a))
		a++
	}
	fmt.Printf("\n")
}
