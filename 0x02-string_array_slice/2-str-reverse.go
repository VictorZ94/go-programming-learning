package main

import "fmt"

func main() {
	str := "Typing nice code is a fuck Art piece!"
	str2 := "Victor Alonso Zuluaga Ramirez"
	reverse(str)
	reverse(str2)
}

func reverse(str string) {
	length := len(str) - 1
	i := 0
	for length >= i {
		fmt.Printf("%c", str[length])
		length--
	}
	fmt.Println()
}
