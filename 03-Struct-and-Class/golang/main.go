package main

import (
	"fmt"
)

type Employee struct {
	Id        int
	FirstName string
	LastName  string
	Age       int
	Department
}

type Department struct {
	Id   int
	Name string
}

func (e *Employee) processPayroll() error {
	fmt.Println("processing payroll")
	// return errors.New("empty name")
	return nil
}

func main() {
	// struct litral
	employee1 := Employee{
		FirstName: "First",
		Age:       20,
		Department: Department{
			Name: "Technology",
		},
	}

	processPayrollError := employee1.processPayroll()
	if processPayrollError != nil {
		fmt.Println("Paylroll processing error")
	}

	fmt.Println(employee1)
	fmt.Println(employee1.FirstName, employee1.Department.Name)

	// dot notation
	var employee2 Employee
	employee2.FirstName = "Second"
	fmt.Println(employee2)

	//new - allocates zeroed storage
	employee3 := new(Employee)
	employee3.FirstName = "Third"
	fmt.Println(employee3)

	var employee4 = new(Employee)
	employee4.FirstName = "Fourth"
	fmt.Println(employee4)

	employee5 := &Employee{}
	employee5.FirstName = "Fifth"
	fmt.Println(employee5.Name)
}
