package main

import "fmt"

func main() {
	var a = 10
	var b = 5
	var sum = a + b
	fmt.Print(sum)

	// In go scope of the variables depend on case of the variable name
	user := "Ram" //the scope of Camel Case variable is inside the package it will not exported to other packages

	UserName := "Ram Jondhale" //the scope of Pascal Case variables  are gloabal it will be exported to other packages
	fmt.Println(user, UserName)

}
