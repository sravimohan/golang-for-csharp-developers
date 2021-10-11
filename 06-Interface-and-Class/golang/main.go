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

func (d *Description) GetName() string {
	return d.Name
}

type Dog struct {
	Description
}

func (d *Dog) Speak() string {
	return "bark"
}

type Cat struct {
	Description
}

func (c *Cat) Speak() string {
	return "meow"
}

func main() {
	rover := Dog{}
	rover.Name = "rover"

	tubby := Cat{}
	tubby.Name = "tubby"

	animals := [2]Animal{&rover, &tubby}
	for _, a := range animals {
		fmt.Printf("%s the %s says %s\n", a.GetName(), reflect.TypeOf(a), a.Speak())
	}
}
