package main

import (
	"fmt"
	"reflect"
)

type Animal interface {
	getName() string
	speak() string
}

// Dog
type Dog struct {
	Name string
}

func (d Dog) getName() string {
	return d.Name
}

func (d Dog) speak() string {
	return "bark"
}

// Cat
type Cat struct {
	Name string
}

func (c Cat) getName() string {
	return c.Name
}

func (c Cat) speak() string {
	return "meow"
}

func getSound(animal Animal) {
	fmt.Println(animal.speak())
}

func main() {
	var rover = Dog{Name: "rover"}
	var tubby = Cat{Name: "tubby"}

	var animals = [2]Animal{rover, tubby}

	for _, a := range animals {
		fmt.Println(a.getName(), " the ", reflect.TypeOf(a), " says ", a.speak())
	}
}
