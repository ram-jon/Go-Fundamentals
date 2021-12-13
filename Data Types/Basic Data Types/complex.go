package main

import (
	"fmt"
)

func main() {
	var a complex128 = complex(6, 2) //complex128 uses float64 for real and imaginary component
	var b complex64 = complex(9, 2)  //complex64 uses float32 for real and imaginary component
	c := 4 + 6i
	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(c)
	// Display the type
	fmt.Printf("The type of a is %T and "+
		"the type of b is %T\n", a, b)

	//extract real and imaginary values
	fmt.Printf("%v, %T\n", real(a), real(a))
	fmt.Printf("%v, %T\n", imag(a), imag(a))

	//all arithmetic operations possible on complex datatype
	fmt.Println(a + c)
}
