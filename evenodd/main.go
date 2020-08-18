package main

import "fmt"

func main() {
	fmt.Println("Starting the even odd program")
	numberSeries := newSeries(0,10)
	numberSeries.checkEvenOdd()
}