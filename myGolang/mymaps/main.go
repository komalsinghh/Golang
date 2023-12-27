package main

import "fmt"

func main(){

	fmt.Println("Maps in Golang")
	languages := make(map[string]string)

	languages["JS"] = "JavaScript"
	languages["RB"] = "Ruby"
	languages["PY"] = "Python"

	fmt.Println("List of languages:",languages)
	fmt.Println("JS shorts for:", languages["JS"])

	delete(languages,"RB")
	fmt.Println("Updated languages:", languages)

	for key, value := range languages{
		fmt.Printf("For key %v, value is %v\n",key,value)
	}
	//or alternative 
	for _, value := range languages{
		fmt.Printf("For key v, value is %v\n",value)
	}
}