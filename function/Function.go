package main

import (
	 "fmt"
	 "strconv"
 )

func main() {
	number := 5
	eNumber := 6

	// calling the function with return type void
	displayFactorial(number)

	// calling the function which return one variable value
	even := evenNumber(eNumber)
	if even {
		fmt.Println(eNumber, "is even number")
	} else {
		fmt.Println(eNumber, "is odd number")
	}

	// calling the function which return two variable values
	add, sub := calc(eNumber, number)
	fmt.Println("Addition of", eNumber, "and", number, "is", add)
	fmt.Println("Substraction of", eNumber, "and", number, "is", sub)
}

// function with return type as void
func displayFactorial(number int) {
	factorial := 1
	equation := "1"
	for index := 2; index <= number; index++ {
		factorial *= index
		equation += " x " + strconv.Itoa(index)
	}
	fmt.Println("Factorial of number", number, "is",equation,"=", factorial)
}

// function return the output
func evenNumber(number int) bool {
	if number%2 == 0 {
		return true
	}
	return false
}

// function to return multiple variable as output
func calc(number1 int, number2 int) (int, int) {
	addition := number1 + number2
	substraction := number1 - number2
	return addition, substraction
}
