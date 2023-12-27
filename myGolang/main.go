package main

import "fmt"

func main(){
	myNumber := 23
	var ptr = &myNumber
	fmt.Println("Value of ptr", ptr)
	fmt.Println("Value of *ptr", *ptr)

	*ptr = *ptr + 2
	fmt.Println("New value", myNumber)

	//how to remove a value from slices based on index
	var courses = []string{"reactjs", "javascript", "swift", "ruby"}
	fmt.Println(courses)
	//remove second index from slice
	var index int = 2
	courses = append(courses[:index], courses[index+1:]...) 
	fmt.Println(courses)
}