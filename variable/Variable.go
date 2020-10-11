package main

import (
	"fmt"
)

func main() {
	// declaring the variable of type int using var keyword
	var num int = 10
	fmt.Println("Value of num is", num)

	// declarin the float variable
	var floatNum float32 = 12.45
	fmt.Println("Value of floatNum is", floatNum)

	// declaring the string variable
	var firstName string = "Manoj"
	fmt.Println("Value of firstName is", firstName)

	// declaring the variable by omitting datatype keyword
	lastName := "Pawar"
	fmt.Println("Value of lastName is", lastName)

	// declaring the constant
	const pi = 3.14
	fmt.Println("Value of pi is", pi)

	// declaring multiple variables
	age, fvtProgLang := 26, "Go"
	fmt.Println("Age is", age, "years")
	fmt.Println("Favorite Programming Language is", fvtProgLang)
}
