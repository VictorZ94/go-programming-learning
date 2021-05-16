package main

import "fmt"

func main() {
	str := "Typing nice code is a fuck Art piece!"
	fmt.Println(length(str))
}

func length(str string) int {
	length := len(str)
	return length
}
