package main

import (
	"errors"
	"fmt"
)

func main() {
	defer CleanUp()
	defer RecoverFromPanic()

	var x, y float32

	fmt.Println("Enter x:")
	fmt.Scanln(&x)

	fmt.Println("Enter y:")
	fmt.Scanln(&y)

	answer, divideError := Divide(x, y)
	if divideError != nil {
		fmt.Println(divideError.Error())
	} else {
		fmt.Println("with error handling: ", x, "/", y, "=", answer)
	}

	answer = DivideWithPanic(x, y)
	fmt.Println("with panic : ", x, "/", y, "=", answer)
}

var DivideByZeroError = errors.New("Cannot Divide by Zero")

func Divide(x float32, y float32) (float32, error) {
	if y == 0 {
		return 0, DivideByZeroError
	}

	return x / y, nil
}

func DivideWithPanic(x float32, y float32) float32 {
	if y == 0 {
		panic("DivideWithPanic")
	}

	return x / y
}

func RecoverFromPanic() {
	p := recover()
	if p != nil {
		fmt.Println("recovering from panic")
	}
}

func CleanUp() {
	fmt.Println("cleaning up")
}
