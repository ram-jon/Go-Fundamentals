package main

import "fmt"

func main() {
	// default value of integer is 0
	var a int32 = -100 //signed int
	var b uint16 = 400 //unsigned int

	//in go both signed and unsigned int available in 4 sizes
	// 8-bit, 16-bit, 32-bit, 64-bit

	// var x,y,x int //Declare multiple variables on same line
	// a,b = 32,42 //assign value to multiple variable

	c := 300 //shorthand for var c = 300

	fmt.Println(a)
	fmt.Println(b)
	fmt.Printf("%v, %T \n", b, b) //we can provide custom format to Printf, %v returns value and %T returns type of variable
	fmt.Println(c)
}
