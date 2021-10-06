package main

import "fmt"

type Person struct {
	Name string
}

func main() {
	// void return
	FunctionWithVoidReturn()

	// string return
	name := FunctionWithStringReturn()
	fmt.Println("FunctionWithStringReturn", name)

	// tuple return
	ok1, value1 := FunctionWithTupleReturn()
	fmt.Println("FunctionWithTupleReturn", ok1, value1)

	// multiple return
	ok2, value2, n2 := FunctionWithMultipleReturn()
	fmt.Println("FunctionWithMultipleReturn", ok2, value2, n2)

	// by value
	name = "Hello"
	FunctionByValue(name)
	fmt.Println("FunctionByValue", name)

	// by reference
	name = "Hello"
	FunctionByReference(&name)
	fmt.Println("FunctionByReference", name)

	// struct by value (default)
	var personByValue = Person{Name: "Hello"}
	FunctionWithStruct(personByValue)
	fmt.Println("FunctionWithStruct", personByValue)

	// struct by refernce
	var personByReference = Person{Name: "Hello"}
	FunctionWithStructByReferebce(&personByReference)
	fmt.Println("FunctionWithStructByReferebce", personByReference)
}

func FunctionWithVoidReturn() {
	// nothing
}

func FunctionWithStringReturn() string {
	return "Hello"
}

func FunctionWithTupleReturn() (bool, string) {
	return true, "Hello"
}

func FunctionWithMultipleReturn() (bool, string, int) {
	return true, "Hello", 10
}

func FunctionByValue(name string) {
	name = "Hello World"
}

func FunctionByReference(name *string) {
	*name = "Hello World"
}

func FunctionWithStruct(p Person) {
	p.Name = "Hello World"
}

func FunctionWithStructByReferebce(p *Person) {
	p.Name = "Hello World"
}
