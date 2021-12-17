// defer keyword is use to postpone the execution to the end of enclosing function

package main

import "fmt"

func main() {
	defer fmt.Println("1st")
	fmt.Println("2st")
	defer fmt.Println("3st") //it will execute before 1st because defer statement execute in LIFO manner
	fmt.Println("4st")
	fmt.Println("5st")

}
