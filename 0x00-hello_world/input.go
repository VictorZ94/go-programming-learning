package main

import "fmt"

func main() {
	var name string
	var age int
	var height float32

	fmt.Println("Type your name: ")
	fmt.Scanf("%s", &name)
	fmt.Println("Now, your age: ")
	fmt.Scanf("%d", &age)
	fmt.Println("Finally, your height: ")
	fmt.Scanf("%f", &height)

	fmt.Printf("Hello %s you're %d years old and %.2f of height\n", name, age, height)
}
