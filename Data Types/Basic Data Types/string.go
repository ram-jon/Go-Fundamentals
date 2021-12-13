package main

import (
	"fmt"
)

func main() {
	var a string = "Welcome"
	str := "Hello"

	fmt.Println(a)
	// Display the length of the string
	fmt.Printf("Length of the string is:%d", len(str))

	// Display the string
	fmt.Printf("\nString is: %s", str)

	// Display the type of str variable
	fmt.Printf("\nType of str is: %T\n", str)

	fmt.Println(string(a[1])) //we can get single character from string using [] it will return ascii value to get actual character typecast it to string

	//convert string into byte array of type uint8
	c := []byte(a)
	fmt.Printf("%v, %T\n", c, c)

	//store character into int32 ascii format
	d := 'A'
	fmt.Printf("%v, %T", d, d)
}
