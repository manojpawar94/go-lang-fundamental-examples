package main

import (
	"errors"
	"fmt"
	"os"
)

func main() {
	filename := "invalid.txt"
	openFile(filename)

	filename = "/Users/manojpawar/go_lang_workspace/fileHandling/sample.txt"
	openFile(filename)

	//receives custom error or nil after trying to open the file
	filename, error := openFileCustomError("invalid.txt")
	if error != nil {
		fmt.Println(error)
	} else {
		fmt.Println("file opened", filename)
	}

}

// error handling 
func openFile(filename string) {
	file, err := os.Open(filename)

	//er will be nil if the file exists else it returns an error object
	if err != nil {
		fmt.Println("Error while opening file", err)
	} else {
		fmt.Println("File opened", file.Name())
	}
}

// custome error handling 
func openFileCustomError(filename string) (string, error) {
	file, err := os.Open(filename)

	//err will be nil if the file exists else it returns an error object
	if err != nil {
		//created a new error object and returns it
		return "", errors.New("Custom error message: File name is wrong")
	}
	return file.Name(), nil
}
