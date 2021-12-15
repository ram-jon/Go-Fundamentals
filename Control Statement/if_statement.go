package main

import "fmt"

func main() {

	// simple if statement
	if true {
		fmt.Println("This is simple if statement")
	}

	//We can also initialise variables and use it in if statement
	if i := 2; i == 2 {
		fmt.Println("if block executed")
	}

	// check item present in map or not using if
	shoppingCart := map[string]int{
		"Keyboard": 1000,
		"Mouse":    500,
		"Laptop":   50000,
	}
	shoppingCart["Monitor"] = 15000

	if _, ok := shoppingCart["Monitor"]; ok {
		fmt.Println("Item present in the shopping cart")
	}

	// else block
	if _, ok := shoppingCart["Mic"]; ok {
		fmt.Println("Item present in the shopping cart")
	} else {
		fmt.Println("Item does not exist in shopping cart")
	}

	//else if block
	if i := 2; i == 3 {
		fmt.Println("This is simple if block")
	} else if i == 2 {
		fmt.Println("This is simple else if block")
	} else {
		fmt.Println("This is simple else block")
	}

	// using logical operators inside if to add multiple conditions

	i := 20
	j := 1

	if i >= 0 && j >= 0 {
		fmt.Println("i and j both are positive")
	}

}
