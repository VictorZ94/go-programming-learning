package main

import (
	"fmt"
	str "strings"
)

var print = fmt.Println

func main() {
	my_str := "test"
	print("Contains:  ", str.Contains(my_str, "es"))
	print("Count:     ", str.Count("test", "t"))
	print("HasPrefix: ", str.HasPrefix("test", "te"))
	print("HasSuffix: ", str.HasSuffix("test", "st"))
	print("Join:      ", str.Join([]string{"a", "b"}, "-"))
	print("Repeat:    ", str.Repeat("a", 5))
	print("Replace:   ", str.Replace("foo", "o", "0", -1))
	print("Replace:   ", str.Replace("foo", "o", "0", 1))
	print("Split:     ", str.Split("a-b-c-d-e", "-"))
	print("ToLower:   ", str.ToLower("TEST"))
	print("ToUpper:   ", str.ToUpper("test"))
	print()

	print("Len: ", len("hello"))
	print("Char:", "hello"[1])         // value in ascii table or
	print("Char:", string("hello"[1])) // char in ascii table
}
