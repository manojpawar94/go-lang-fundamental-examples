package main

import (
	"fmt"
	"time"
)

func main() {
	//creating a channel variable to transport integer values
	ch := make(chan string)

	//invoking the subroutines to add and fetch from the channel
	//These routines execute simultaneously
	go writeToChannel(ch)
	go readFromChannel(ch)

	//delay is to prevent the exiting of main() before goroutines finish
	time.Sleep(5 * time.Second)
	fmt.Println("Inside main()")
}

func writeToChannel(ch chan string) {
	fmt.Println("write data to channel...")
	fruits := [...]string{"Apple", "Mango", "Banana", "Grapes", "PineApple"}
	for _, fruit := range fruits {
		ch <- fruit // writing/pushing data to channel
	}
	close(ch)
}

func readFromChannel(ch chan string) {
	fmt.Println("Reading data from channel")
	for {
		// fetch data from channel
		fruit, flag := <-ch

		if flag {
			fmt.Println(fruit)
		} else {
			fmt.Println("Channel is empty")
			break
		}
	}
}
