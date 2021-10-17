package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	a, _ := strconv.Atoi(os.Args[1])
	b, _ := strconv.Atoi(os.Args[2])

	var cs CalculatorService = new(Calculator)

	fmt.Println(a, "+", b, "=", cs.Add(a, b))
	fmt.Println(a, "-", b, "=", cs.Subtract(a, b))
	fmt.Println(a, "/", b, "=", cs.Divide(a, b))
	fmt.Println(a, "*", b, "=", cs.Multiply(a, b))
}
