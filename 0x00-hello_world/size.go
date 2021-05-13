package main

import (
	"fmt"
	"unsafe"
)

func main() {
	boolean := bool(true)
	fmt.Printf("Size of a bool: %d byte(s)\n", unsafe.Sizeof(boolean))

	str := string("H")
	fmt.Printf("Size of a string: %d byte(s)\n", unsafe.Sizeof(str))

	num := int(1)
	fmt.Printf("Size of a int: %d byte(s)\n", unsafe.Sizeof(num))

	num1 := int8(1)
	fmt.Printf("Size of a int8: %d byte(s)\n", unsafe.Sizeof(num1))

	num2 := int16(1)
	fmt.Printf("Size of a int16: %d byte(s)\n", unsafe.Sizeof(num2))

	num3 := int32(1)
	fmt.Printf("Size of a int32: %d byte(s)\n", unsafe.Sizeof(num3))

	num4 := int64(1)
	fmt.Printf("Size of a int64: %d byte(s)\n", unsafe.Sizeof(num4))

	unum := int(1)
	fmt.Printf("Size of a uint: %d byte(s)\n", unsafe.Sizeof(unum))

	unum1 := int8(1)
	fmt.Printf("Size of a uint8: %d byte(s)\n", unsafe.Sizeof(unum1))

	unum2 := int16(1)
	fmt.Printf("Size of a uint16: %d byte(s)\n", unsafe.Sizeof(unum2))

	unum3 := int32(1)
	fmt.Printf("Size of a uint32: %d byte(s)\n", unsafe.Sizeof(unum3))

	unum4 := int64(1)
	fmt.Printf("Size of a uint64: %d byte(s)\n", unsafe.Sizeof(unum4))

	ptr := uintptr(1)
	fmt.Printf("Size of a uintptr: %d byte(s)\n", unsafe.Sizeof(ptr))

	floatant := float32(1.5)
	fmt.Printf("Size of a float32: %d byte(s)\n", unsafe.Sizeof(floatant))

	floatant2 := float64(1.5)
	fmt.Printf("Size of a float64: %d byte(s)\n", unsafe.Sizeof(floatant2))

	// byte type is an alias of int8
	char := byte(86)
	fmt.Printf("Size of a byte: %d byte(s) %c \n", unsafe.Sizeof(char), char)

	// rune type is an alias of int32
	runes := rune(0xC6)
	fmt.Printf("Size of a rune: %d byte(s), %c\n", unsafe.Sizeof(runes), runes)

	numcomplex := complex64(1)
	fmt.Printf("Size of a complex64: %d byte(s)\n", unsafe.Sizeof(numcomplex))

	numcomplex2 := complex128(1)
	fmt.Printf("Size of a complex128: %d byte(s)\n", unsafe.Sizeof(numcomplex2))

	// fmt.Printf("a: %T, %d, %d\n", num, unsafe.Sizeof(num), num)
}
