package main

import (
	"fmt"
	"unsafe"
)

func main() {
	var a = "initial"
	fmt.Println(a)

	var b, c int = 1, 2
	fmt.Println(b, c)

	var height = 1.83
	fmt.Println(height)

	const pi = 1.3433
	fmt.Println(pi)

	var d = true
	fmt.Println(d)

	num := int64(123)
	fmt.Printf("a: %T, %d, %d\n", num, unsafe.Sizeof(num), num)

	var e int
	fmt.Println(e)

	f := "apple"
	fmt.Println(f)

	//Relational operators
	/*
	 >
	 <
	 >=
	 <=
	 ==
	 all return Boolean value
	*/

	//Logic operators in go
	/*
	 &&
	 ||
	 !
	*/
}
