package main

import "fmt"

func main() {
	num := 0
	base16 := 65
	for num < 16 {

		if num > 9 {
			fmt.Printf("%c", byte(base16))
			base16++
		} else {
			fmt.Printf("%d", num)
		}
		if num < 15 {
			fmt.Printf(", ")
		}
		num++

	}
	fmt.Printf("\n")
}
