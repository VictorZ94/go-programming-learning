package main

import "fmt"

func main() {
	num := 0
	for num <= 9 {
		fmt.Printf("%d", num)
		if num < 9 {
			fmt.Printf(", ")
		}
		num++
	}
	fmt.Printf("\n")
}
