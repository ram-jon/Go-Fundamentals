package main

import (
	"fmt"
	"strconv"
)

func main() {
	var a int = 90
	var b float64 = float64(a)
	var c string = string(a)       //typecasting integer to string result in ascii character value
	var d string = strconv.Itoa(a) //Itoa method of strconv package convert integer to string digits
	fmt.Printf("%v, %T\n", a, a)
	fmt.Printf("%v, %T\n", b, b)
	fmt.Printf("%v, %T\n", c, c)
	fmt.Printf("%v, %T\n", d, d)

}
