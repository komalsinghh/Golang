package main

import (
	"fmt"
	"net/url"
)


const myurl string ="https://lco.dev:3000/learn?coursename=reactjs&paymentid=ghbj456ghb"

func main(){
	fmt.Println(myurl)

	//parsing url
	result, _ :=url.Parse(myurl)
	fmt.Println(result.Scheme)
	fmt.Println(result.Host)
	fmt.Println(result.Path)
	fmt.Println(result.Port())
	fmt.Println(result.RawQuery)

	qparams := result.Query()
	fmt.Printf("The type of query params are: %T\n",qparams)

	fmt.Println(qparams["paymentid"])

	for _, val:= range qparams{
		fmt.Println("Param is: ",val)
	}

	//constructing an url
	partsOfUrl := &url.URL{
		Scheme: "https",
		Host: "loc.dev",
		Path: "/tutcss",
		RawPath: "user=hitesh",
	}

	anotherURL := partsOfUrl.String()
	fmt.Println(anotherURL)
}	