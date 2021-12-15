// Anonymous functions are functions without name
// it is useful when we need to create inline function
//we can also pass  and return anonymous funtion  from another function

package main

import "fmt"

func main() {
	// create and call anonymous function
	func(msg string) {
		fmt.Println(msg)
	}("Hello World")

	// create and assign anonymous function to variable
	fun := func(a int, b int) int {
		return a + b
	}
	fmt.Println(fun(5, 2)) //call function
}
