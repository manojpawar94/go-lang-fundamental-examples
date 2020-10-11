package main

import (
	"fmt"
	"io/ioutil"
)

func readFile(filename string) {
	data, err := ioutil.ReadFile(filename)

	if err != nil {
		fmt.Println("File reading error", err)
		return
	}

	fmt.Println("Contents of file:", string(data))
}
