package main

import (
	"fmt"
	"time"
)

func main() {
	value := 5
	
	// creating the channel to communicate between routines
	ch := make(chan bool)

	// passing channel to routine
	go counter(value, ch)
	// read data from the routine
	status := <-ch

	// closing the channel
	close(ch)

	fmt.Println("Inside main()")
	fmt.Println("Printing status in main() after taking from channel:", status)
}

func counter(value int, ch chan bool) {
	fmt.Println("Starting counter...")
	for index := 0; index < value; index++ {
		time.Sleep(1 * time.Second)
	}
	fmt.Println("Completed counter...")
	// passing data to channel 
	ch <- true
}
