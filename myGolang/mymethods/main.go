package main

import "fmt"

func main(){
	komal := User{"Komal", "komal@go.dev", true, 23}
	fmt.Println(komal)
	fmt.Printf("komal details are: %+v\n",komal)
	fmt.Printf("Name is %v and email is %v.\n",komal.Name,komal.Email)
	komal.GetStatus()
	komal.NewMail()
}

type User struct{
	Name string
	Email string
	Status bool
	Age int
}

func (u User) GetStatus(){
	fmt.Println("Is user active", u.Status)
}

func (u User) NewMail(){
	u.Email="test@go.dev"
	fmt.Println("Email of this user is", u.Email)
}