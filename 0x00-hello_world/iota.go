package main

import "fmt"

// const Sunday int = 0
// const Monday int = 1
// const Tuesday int = 2
// const Wednesday int = 3
// const Thursday int = 4
// const Friday int = 5
// const Saturday int = 6

// this is the same
// const (
// 	Sunday    int = 0
// 	Monday    int = 1
// 	Tuesday   int = 2
// 	Wednesday int = 3
// 	Thursday  int = 4
// 	Friday    int = 5
// 	Saturday  int = 6
// )

const (
	Sunday int = iota + 1 // initialize in 0. it indicates is sequentially values
	Monday
	Tuesday
	Wednesday
	Thursday
	Friday
	Saturday
)

func main() {

	fmt.Println(Sunday)
	fmt.Println(Monday)
	fmt.Println(Tuesday)
	fmt.Println(Wednesday)
	fmt.Println(Thursday)
	fmt.Println(Friday)
	fmt.Println(Saturday)
}
