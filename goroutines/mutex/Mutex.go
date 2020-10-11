package main

import (
	"fmt"
	"math/rand"
	"strconv"
	"sync"
	"time"
)

// declaring the global Mutex instance
var mu sync.Mutex

// declaring the global count variable
var count = 0

func main() {
	for index := 1; index < 4; index++ {
		go dummyProcess(index)
	}

	//delay to wait for the routines to complete
	time.Sleep(25 * time.Second)
	fmt.Println("Final Count:", count)
}

func dummyProcess(index int) {
	for itr := 0; itr < 10; itr++ {
		time.Sleep(time.Duration(rand.Int31n(2)) * time.Second)

		// sync block start
		mu.Lock()
		temp := count
		temp++
		time.Sleep(time.Duration(rand.Int31n(2)) * time.Second)
		count = temp
		mu.Unlock()
		// sync block end
	}
	fmt.Println("Count after index=", strconv.Itoa(index), " Count=", count)
}
