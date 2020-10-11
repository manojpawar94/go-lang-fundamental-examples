package main

import "fmt"

func main() {
	// declare an array
	var evenNumbers [5]int
	evenNumbers[0] = 0
	evenNumbers[1] = 2
	evenNumbers[2] = 4
	evenNumbers[3] = 6
	evenNumbers[4] = 8

	// accessing the element of the array using the index
	fmt.Println("Even number at position 3 in evenNumbers is", evenNumbers[2])
	fmt.Println("Even number at index 3 in evenNumbers is", evenNumbers[3])

	// calculating the length of the array
	fmt.Println("The length of evenNumbers is", len(evenNumbers))

	// declaring array and assign value in same step
	seasons := [3]string{"Summar", "Winter", "Rain"}

	fmt.Println("Printing seasons,")
	// iterating the array using for loop
	for index, season := range seasons {
		fmt.Println(index, ":", season)
	}

	// declaring array and assign value in same step
	fruits := [...]string{"Apple", "Mango", "Banana", "PineApple"}

	// printing array
	fmt.Println("Fruits :", fruits)
}
