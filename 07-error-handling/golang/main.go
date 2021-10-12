package main

import (
	"errors"
	"fmt"
)

func main() {
	defer CleanUp()
	defer RecoverFromPanic()

	var x float32
	var y float32

	fmt.Print("Enter x:")
	fmt.Scanln(&x)

	fmt.Print("Enter y:")
	fmt.Scanln(&y)

	answer, divideError := Divide(x, y)

	if divideError == nil {
		fmt.Println(x, "/", y, "=", answer)
	} else {
		fmt.Println(divideError.Error())
	}

	answer = DivideWithPanic(x, y)
}

var DivideByZero = errors.New("Cannot divide by Zero")

func Divide(x float32, y float32) (float32, error) {
	if y == 0 {
		return 0, DivideByZero
	}

	return x / y, nil
}

func DivideWithPanic(x float32, y float32) float32 {
	if y == 0 {
		panic("DivideWithPanic() y is zero")
	}

	return x / y
}

func CleanUp() {
	fmt.Println("Cleaning Up")
}

func RecoverFromPanic() {
	p := recover()
	if p != nil {
		fmt.Println("main(): recover", p)
	}
}
