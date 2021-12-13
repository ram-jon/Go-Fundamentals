package main

import "fmt"

func main() {
	var p int = 23
	var q int = 60
	//Logical AND
	if p != q && p <= q {
		fmt.Println("True")
	}

	// Logical OR
	if p != q || p <= q {
		fmt.Println("True")
	}

	// Logical NOT
	if !(p == q) {
		fmt.Println("True")
	}

}
