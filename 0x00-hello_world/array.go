package main

import "fmt"

func main() {

	var a [5]int
	a[0] = 100
	a[1] = 200
	a[2] = 200
	a[3] = 200
	a[4] = 200
	fmt.Println(a)

	//another form
	num := [5]int{1, 2, 3, 4, 5}
	fmt.Println(num)
	//another form without indicating its size
	b := [...]int{10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20}
	fmt.Println(b)

	//I also can indicate the index in the declared array
	arr := [...]int{2: 12, 0: 13, 1: 14}
	fmt.Println(arr)

	currency := [...]string{2: "COP", 0: "USD", 1: "MPE", 5: "SOL"}
	fmt.Println(currency)

	//I also can create subarrays ---===---- this are slice
	subCurrency := currency[0:2]
	fmt.Println(subCurrency)
}
