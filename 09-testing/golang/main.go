package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	a, _ := strconv.Atoi(os.Args[1])
	b, _ := strconv.Atoi(os.Args[2])

	c := Calculator{}
	fmt.Println(a, "+", b, "=", c.Add(a, b))
	fmt.Println(a, "+", b, "=", c.Subtract(a, b))
	fmt.Println(a, "+", b, "=", c.Multiply(a, b))
	fmt.Println(a, "+", b, "=", c.Divide(a, b))
}
