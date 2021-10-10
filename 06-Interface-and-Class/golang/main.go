package main

import (
	"fmt"
	"reflect"
)

type Animal interface {
	GetName() string
	Speak() string
}

type Description struct {
	Name string
}

func (a Description) GetName() string {
	return a.Name
}

// Dog
type Dog struct {
	Description
}

func (a *Dog) Speak() string {
	return "bark"
}

// Cat
type Cat struct {
	Description
}

func (c *Cat) Speak() string {
	return "meow"
}

func getSound(animal Animal) {
	fmt.Println(animal.Speak())
}

func main() {
	rover := Dog{}
	rover.Name = "rover"

	tubby := Cat{}
	tubby.Name = "tubby"

	var animals = [2]Animal{&rover, &tubby}
	for _, a := range animals {
		fmt.Printf("%s the %s says %s\n", a.GetName(), reflect.TypeOf(a), a.Speak())
	}
}
