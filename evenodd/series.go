package main

import "fmt"

type intSeries []int

// creating the number series with default step size 1
func newSeries(start int, end int) intSeries {
	return newSeriesStep(start, end, 1)
}

// creating the number series with step
func newSeriesStep(start int, end int, step int) intSeries {
	series := intSeries{}

	for index := start; index <= end; index += step {
		series = append(series, index)
	}
	return series
}

// check and print numbers from slice even or odd
func (series intSeries) checkEvenOdd() {
	for _,num := range series {
		if num%2 == 0 {
			fmt.Println(num, "is even")
		} else {
			fmt.Println(num, "is odd")
		}
	}
}
