package main

import "fmt"

/*Employee define the struct of employee*/
type Employee struct {
	id         int
	name       string
	department string
}

// function accept variable of type employee struct
func display(employee Employee) {
	fmt.Println("Employee( id :", employee.id, ", name :", employee.name, ", department :", employee.department, ")")
}

//Declaring a function with receiver of the type employee
func (employee Employee) toString() {
	fmt.Println("Employee( id :", employee.id, ", name :", employee.name, ", department :", employee.department, ")")
}
