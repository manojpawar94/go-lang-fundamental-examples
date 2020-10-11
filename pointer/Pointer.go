package main

import "fmt"

func main() {
	// declaring the variable
	name := "Manoj"
	namePointer := &name

	fmt.Println("name =", name)
	// printing the pointer variable of name
	fmt.Println("namePointer =", namePointer)
	// print the value at using pointer variable
	fmt.Println("*namePointer =", *namePointer)

	// modifying the value using pointer variable
	*namePointer = *namePointer + " Pawar"
	fmt.Println("name =", name)
	fmt.Println("namePointer =", namePointer)
	fmt.Println("*namePointer =", *namePointer)

}
