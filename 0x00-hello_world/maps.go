package main

import "fmt"

func main() {
	day := make(map[int]string)

	day[0] = "Sunday"
	day[1] = "Monday"
	day[2] = "Tuesday"
	day[3] = "Wednesday"
	day[4] = "Tuesday"
	day[5] = "friday"
	day[6] = "Saturday"

	day[4] = "Jueves"
	delete(day, 4)

	fmt.Println(day)

	//structe of map more complex
	user := make(map[string][]int) // Dict of arrays

	user["edwart"] = []int{1, 2, 3, 4, 5}
	fmt.Println(user["edwart"][3])
	fmt.Println(user["edwart"])
	fmt.Println(user)

	//iter an map
	my_map := map[int]string{} // another form to declare an map
	my_map[3] = "Victor"
	my_map[0] = "Deisy"
	my_map[1] = "Octavio"
	my_map[2] = "Elvia"

	for key, value := range my_map {
		fmt.Println(key, value)
	}
}
