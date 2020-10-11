package main

import "fmt"

func main() {
	// creating the slice from array
	numArray := [...]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 0}
	sliceArrayOne := numArray[2:5]
	sliceArrayTwo := numArray[5:9]

	// printing slice
	fmt.Println("numArray is", numArray)
	fmt.Println("sliceArrayOne is", sliceArrayOne)
	fmt.Println("sliceArrayOne is", sliceArrayTwo)

	// calculating length of slice
	fmt.Println("Length of numArray is", len(numArray))
	fmt.Println("Length of sliceArrayOne is", len(sliceArrayOne))
	fmt.Println("Length of sliceArrayTwo is", len(sliceArrayTwo))

	// calculating capacity of slice
	fmt.Println("Capacity of numArray is", cap(numArray))
	fmt.Println("Capacity of sliceArrayOne is", cap(sliceArrayOne))
	fmt.Println("Capacity of sliceArrayTwo is", cap(sliceArrayTwo))

	//appending one slice to another
	newSliceAfterAppend := append(sliceArrayOne, sliceArrayTwo...)
	fmt.Println("New Slice :", newSliceAfterAppend)

	//append two slice into single slice
	sliceArrayTwo = append(sliceArrayTwo, sliceArrayOne...)
	fmt.Println("Updated sliceArrayTwo :", newSliceAfterAppend)

	//append elment to the slice
	sliceArrayOne = append(sliceArrayOne, 12)
	fmt.Println("Newly appended sliceArrayOne :", sliceArrayOne)
}
