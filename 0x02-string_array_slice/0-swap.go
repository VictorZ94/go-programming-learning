package main

import "fmt"

func main() {
	a := 5
	b := 10
	fmt.Println("a is:", a, "and b is:", b)
	b, a = a, b
	fmt.Println("a is:", a, "and b is:", b)
}
