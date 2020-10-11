package main

import "fmt"

func main() {
	// create variable of type employee struct
	var employee1 Employee

	// assign value to the member of struct
	employee1.id = 1
	employee1.name = "Manoj"
	employee1.department = "Technology"

	// prints the members of employee using display function
	display(employee1)

	// declares and assign values to variable employee2 of type employee
	employee2 := Employee{2, "Rajesh", "HR"}

	// Invoking the method using the receiver of the type employee
	// syntax is variable.methodname()
	employee2.toString()

	fmt.Println("Example on init struct")
	example1()

	fmt.Println("Example on nested structs")
	example2()

}
