package main

import (
	"fmt"
	"reflect"
)

type Animal interface {
	getName() string
	speak() string
}

type Description struct {
	Name string
}

func (a Description) getName() string {
	return a.Name
}

// Dog
type Dog struct {
	Description
}

func (a *Dog) speak() string {
	return "bark"
}

// Cat
type Cat struct {
	Description
}

func (c *Cat) speak() string {
	return "meow"
}

func getSound(animal Animal) {
	fmt.Println(animal.speak())
}

func main() {
	var rover = Dog{Description: Description{Name: "rower"}}
	var tubby = Cat{Description: Description{Name: "tubby"}}

	var animals = [2]Animal{&rover, &tubby}
	for _, a := range animals {
		fmt.Printf("%s the %s says %s\n", a.getName(), reflect.TypeOf(a), a.speak())
	}
}
