package main

import "fmt"

/*
Types
	string

	bool

	Numeric Types
		int8, int16, int32, int64, int
		uint8, uint16, uint32, uint64, uint
		float32, float64
		complex64, complex128
			example: 7 + 2i
		byte - alias of uint8
		rune - alias of int32
*/

func main() {
	var id int = 1
	var name = "hello"
	age := 20
	fmt.Println(id, name, age)

	num1 := 10
	num2 := 10.5
	total := float64(num1) + num2
	fmt.Println(num1, num2, total)

	var c1 complex64 = complex(1, 2)
	fmt.Println(c1)

	// '♛', '♠', '♧', '♡'
	symbols := []rune{'♛', '♠', '♧', '♡'}
	fmt.Println(symbols)
}
