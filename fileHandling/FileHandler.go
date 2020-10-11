package main

func main() {
	filename := "sample.txt"
	content := "This sample text file, to test create and read file"

	writeFile(filename, content)
	readFile(filename)
}
