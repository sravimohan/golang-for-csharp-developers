package main

import (
	"testing"
)

func TestAdd(t *testing.T) {
	a := 10
	b := 2
	expected := 12

	c := Calculator{}
	actual := c.Add(a, b)

	if expected != actual {
		t.Errorf("Add was incorrect, expected: %d, expected: %d.", expected, actual)
	}
}

func TestSubtract(t *testing.T) {
	a := 10
	b := 2
	expected := 8

	c := Calculator{}
	actual := c.Subtract(a, b)

	if expected != actual {
		t.Errorf("Subtract was incorrect, expected: %d, expected: %d.", expected, actual)
	}
}

func TestMultiply(t *testing.T) {
	a := 10
	b := 2
	expected := 20

	c := Calculator{}
	actual := c.Multiply(a, b)

	if expected != actual {
		t.Errorf("Multiply was incorrect, expected: %d, expected: %d.", expected, actual)
	}
}

func TestDivide(t *testing.T) {
	a := 10
	b := 2
	expected := 5

	c := Calculator{}
	actual := c.Divide(a, b)

	if expected != actual {
		t.Errorf("Divide was incorrect, expected: %d, expected: %d.", expected, actual)
	}
}
