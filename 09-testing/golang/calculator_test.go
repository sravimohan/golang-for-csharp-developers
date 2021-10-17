package main

import "testing"

func TestAdd(t *testing.T) {
	a := 10
	b := 2
	expected := 12

	var cs CalculatorService = new(Calculator)
	actual := cs.Add(a, b)

	if expected != actual {
		t.Errorf("Add was incorrect. expected %d; actual: %d", expected, actual)
	}
}

func TestSubtract(t *testing.T) {
	a := 10
	b := 2
	expected := 8

	var cs CalculatorService = new(Calculator)
	actual := cs.Subtract(a, b)

	if expected != actual {
		t.Errorf("Subtract was incorrect. expected %d; actual: %d", expected, actual)
	}
}

func TestDivide(t *testing.T) {
	a := 10
	b := 2
	expected := 5

	var cs CalculatorService = new(Calculator)
	actual := cs.Divide(a, b)

	if expected != actual {
		t.Errorf("Divide was incorrect. expected %d; actual: %d", expected, actual)
	}
}

func TestMultiply(t *testing.T) {
	a := 10
	b := 2
	expected := 20

	var cs CalculatorService = new(Calculator)
	actual := cs.Multiply(a, b)

	if expected != actual {
		t.Errorf("Multiply was incorrect. expected %d; actual: %d", expected, actual)
	}
}
