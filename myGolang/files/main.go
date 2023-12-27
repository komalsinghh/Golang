package main

import (
	"fmt"
	"io"
	"os"
)

func main() {
	fmt.Println("Welcome to files in golang")
	content := "This needs to go in a file - golang example"

	file, err := os.Create("./mygolangfile.txt")
	checkNilErr(err)
	// if err != nil {
	// 	panic(err)
	// }

	length, err := io.WriteString(file, content)
	checkNilErr(err)
	fmt.Println("length is:",length)
	defer file.Close()
	readFile("./mygolangfile.txt")
}

func readFile(fileName string){
	databyte, err := os.ReadFile(fileName)
	checkNilErr(err)

	fmt.Println("Text data inside the file is \n", string(databyte))
}

func checkNilErr(err error){
	if err != nil {
		panic(err)
	}
}