package main

import (
	"fmt"
)

func main() {
	num := []int{1, 2, 3, 4, 5} // slices are like reference to an array base
	num = append(num, 6)
	num = append(num, 7)
	num = append(num, 8)
	num = append(num, 9)
	num = append(num, 10)
	fmt.Println(num)

	new := num[0:5]
	new[0] = 100
	fmt.Println(num)
	fmt.Println(new)

	//Slice of string
	month := []string{
		"january",
		"February",
		"March",
		"April",
		"May",
		"June",
		"July",
		"August",
		"September",
	}

	//Cuando hablamos de slice debemos pensar en lo siguiente
	/*
		Un puntero or memory address
		Una longitud
		Una Capacidad
	*/

	length := len(month)
	capac := cap(month)

	fmt.Printf("La longitud es: %v \nLa capacidad es: %v \nmemory address: %p\n", length, capac, month)

	month = append(month, "October") // Si la estructura se encuentra en su máxima capacidad, se crea una nueva estructura
	month = append(month, "November")
	month = append(month, "December")
	length = len(month)
	capac = cap(month)

	fmt.Printf("La longitud es: %v \nLa capacidad es: %v \nmemory address: %p\n", length, capac, month)

}
