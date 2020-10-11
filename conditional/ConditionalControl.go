package main

import "fmt"

func main() {
	// if else statement
	const ageLimit = 18
	age := 22
	if age < ageLimit {
		fmt.Println("Age", age, "years is lower than the age limit of", ageLimit, "years.")
	} else {
		fmt.Println("Age", age, "years is greater than or equal to the age limit of", ageLimit, "years.")
	}

	// nested if else statement
	num := 100
	if num < 10 {
		//Executes if num is less than 10
		fmt.Println("Num is less than 10")
	} else if num >= 10 && num <= 90 {
		//Executes if num >= 10 and num<=90
		fmt.Println("Num is between 10 and 90")
	} else {
		//Executes if both above cases fail i.e num>90
		fmt.Println("Num is greater than 90")
	}

	// switch case
	colorCode := 1
	var colorValue string
	switch colorCode {
	case 1:
		colorValue = "Blue"
	case 2:
		colorValue = "Red"
	case 3:
		colorValue = "Green"
	default:
		colorValue = "White"
	}
	fmt.Println("The color code", colorCode, "is", colorValue, "color")
}
