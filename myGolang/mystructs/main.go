package main

import "fmt"

func main(){
	komal := User{"Komal", "komal@go.dev", true, 23}
	fmt.Println(komal)
	fmt.Printf("komal details are: %+v\n",komal)
	fmt.Printf("Name is %v and email is %v.\n",komal.Name,komal.Email)
}

type User struct{
	Name string
	Email string
	Status bool
	Age int
}