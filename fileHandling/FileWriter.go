package main

import (
	"fmt"
	"os"
)

func writeFile(filename string, content string) {
	file, err := os.Create(filename)

	if err != nil {
		fmt.Println("Error while creating file", err)
		return
	}

	l, err := file.WriteString(content)

	if err != nil {
		fmt.Println("Error while writting content to file", err)
		file.Close()
		return
	}

	fmt.Println(l, "Bytes written")

	err = file.Close()
	if err != nil {
		fmt.Println("Error while closing file",err)
		return
	}
}
