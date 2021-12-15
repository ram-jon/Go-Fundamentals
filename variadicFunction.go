package main

import (
	"fmt"
	"strings"
)

// The function that called with the varying number of arguments is known as variadic function. Or in other words, a user is allowed to pass zero or more arguments in the variadic function.
func main() {
	fmt.Println(sum(1, 2, 3, 4, 5))
	sumMsg("Hello", 1, 2, 3, 4, 5)

	fmt.Println(joinstr("Hello", "Welcome", "to", "Go"))

	// we can also pass a slice in variadic function
	element := []string{"Hello", "Welcome", "to", "Go"}
	fmt.Println(joinstr(element...))
}

func sum(values ...int) int {
	sum := 0
	for _, v := range values {
		sum += v
	}
	return sum
}

func sumMsg(msg string, values ...int) { //only last parameter can be of type variadic
	sum := 0
	for _, v := range values {
		sum += v
	}
	fmt.Println(msg, sum)
}

// common usecase

// Variadic function to join strings
func joinstr(element ...string) string {
	return strings.Join(element, "-")
}
