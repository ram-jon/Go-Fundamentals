package main

import "fmt"

func main() {

	switch 2 {
	case 1:
		fmt.Println("This is 1")
	case 2:
		fmt.Println("This is 2")
	default:
		fmt.Println("this is default")
	}

	// Adiing multiple case conditions using ","

	switch i := 2 + 3; i {
	case 1, 3, 5, 7, 9:
		fmt.Println("This is Odd")
	case 2, 4, 6, 8:
		fmt.Println("This is Even")
	default:
		fmt.Println("This is Default")
	}

	//Tagless switch statement
	i := 2 + 3
	switch {
	case i > 0:
		fmt.Println("This is positive")
	case i < 0:
		fmt.Println("This is negative")
	default:
		fmt.Println("This is zero")
	}

	// Type switch

	var a interface{} = 5.2
	switch a.(type) {
	case int:
		fmt.Println("This is int type")
	case float64:
		fmt.Println("This is float64 type")
	case string:
		fmt.Println("This is string type")
	default:
		fmt.Println("This is default type")
	}

	// break and fallthrough
	// there is implicit break after every case in go
	switch 1 {
	case 1:
		fmt.Println("This is 1")
		fallthrough //it continue to execute next case statement
	case 2:
		fmt.Println("This is 2")
		fmt.Println("This is 2")
		break //it will stop execution of further statement after break
		fmt.Println("This is 2")
		fmt.Println("This is 2")
	default:
		fmt.Println("this is default")
	}

}
