package main

import "fmt"

func main(){
	proRes := proAdder(2,5,8,7)
	fmt.Println("Pro result is", proRes)
}

func proAdder(values ...int) int{
	total :=0
	for _,value := range values{
		total+=value
	}
	return total
}