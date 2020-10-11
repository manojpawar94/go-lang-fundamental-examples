package main

import (
	"fmt"
	"time"
)

func main() {
	counterValue := 10
	go downCounter(counterValue)
	go counter(counterValue)
	time.Sleep(10 * time.Second)
	fmt.Println("Completed Main Function...")
}

func downCounter(counterValue int) {
	for index := counterValue; index > 0; index-- {
		time.Sleep(1 * time.Second)
		fmt.Println("Down Counter:", index)
	}
	fmt.Println("Down counter completed")
}

func counter(counterValue int) {
	for index := 0; index < counterValue; index++ {
		time.Sleep(1 * time.Second)
		fmt.Println("Counter:", index)
	}
	fmt.Println("Counter completed")
}
