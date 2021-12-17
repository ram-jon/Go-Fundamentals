package main

import "fmt"

func main() {
	writeMessage("Hello")
	num := 5
	fmt.Println(add(4, 5))
	fmt.Println(num)
	passByValue(num)
	fmt.Println(num)
	passByRef(&num)
	fmt.Println(num)

	fmt.Println("---------------------------------")
	output, error := divide(5, 0)
	if error != nil {
		fmt.Println(error)
		return
	}
	fmt.Println(output)
}

// simple function with 1 parameter
func writeMessage(message string) {
	fmt.Println(message)
}

// multiple parameters ans return type int

func add(a, b int) int {
	return a + b
}

// pass by value parameter
func passByValue(num int) {
	num = 10
	fmt.Println(num)
}

// pass by value reference
func passByRef(num *int) {
	*num = 20
	fmt.Println(*num)
}

// returning multiple values

func divide(a float64, b float64) (float64, error) {
	if b == 0 || b == 0.0 { //handling divide by zero
		return 0.0, fmt.Errorf("Cannot divide by zero")
	}
	return a / b, nil
}
