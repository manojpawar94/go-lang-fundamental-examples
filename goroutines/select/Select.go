package main

import (
	"fmt"
	"time"
)

func main() {
	upCounterValue := 3
	downCounterValue := 5
	upCountChan := make(chan int)
	downCountChan := make(chan int)

	go downCounter(downCountChan, downCounterValue)
	go counter(upCountChan, upCounterValue)

	select {
	case value := <-upCountChan:
		fmt.Println("Up counter value", value)
	case value := <-downCountChan:
		fmt.Println("Down counter value", value)
	}

	upCounterValue = 5
	downCounterValue = 3

	select {
	case value := <-upCountChan:
		fmt.Println("Up counter value", value)
	case value := <-downCountChan:
		fmt.Println("Down counter value", value)
	}

	select {
	case value := <-upCountChan:
		fmt.Println("Up counter value", value)
	case value := <-downCountChan:
		fmt.Println("Down counter value", value)
	default:
		fmt.Println("Executing the default case")
	}

	close(upCountChan)
	close(downCountChan)
}

func counter(upCountChan chan int, value int) {
	for index := 0; index < value; index++ {
		time.Sleep(1 * time.Second)
	}
	upCountChan <- value
}

func downCounter(downCountChan chan int, value int) {
	for ; value > 0; value-- {
		time.Sleep(1 * time.Second)
	}
	downCountChan <- value
}
