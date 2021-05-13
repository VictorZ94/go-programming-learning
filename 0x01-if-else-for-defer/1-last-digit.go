package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	rand.Seed(time.Now().UnixNano())
	MAX_RAND := 1000000000
	n := rand.Intn(MAX_RAND) / 2

	mod := n % 10
	if mod > 5 {
		fmt.Printf("Last digit of %d is %d and is greater than 5\n", n, mod)
	} else if mod == 0 {
		fmt.Printf("Last digit of %d is %d and is 0\n", n, mod)
	} else {
		fmt.Printf("Last digit of %d is %d and is less than 6 and not 0\n", n, mod)
	}
}
