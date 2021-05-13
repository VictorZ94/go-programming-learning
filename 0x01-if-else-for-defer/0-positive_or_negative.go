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
	if n > 0 {
		fmt.Printf("%d is Positive\n", n)
	} else {
		fmt.Println(n, " is Negative")
	}
}
