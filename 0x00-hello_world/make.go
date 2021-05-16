package main

import "fmt"

func main() {
	num := make([]int, 3, 3)

	num[0] = 102
	num[1] = 9
	num[2] = 56

	num = append(num, 99) // its cap it duplicates
	fmt.Println(num)
	fmt.Println(len(num))
	fmt.Println(cap(num))
}
