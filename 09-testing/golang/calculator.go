package main

type CalculatorService interface {
	Add(a int, b int) int
	Subtract(a int, b int) int
	Multiply(a int, b int) int
	Divide(a int, b int) int
}

type Calculator struct{}

func (c *Calculator) Add(a int, b int) int {
	return a + b
}

func (c *Calculator) Subtract(a int, b int) int {
	return a - b
}

func (c *Calculator) Multiply(a int, b int) int {
	return a * b
}

func (c *Calculator) Divide(a int, b int) int {
	return a / b
}
