package main

import "fmt"

const user string = "Admin" //global constant available in the package

// defining multiple constant in constant block
const (
	a int = 12
	b int = 13
	c int = 14
)

// inumerated constants  -> auto increment constant value start from 0
const (
	x = iota //0
	y = iota //1
	z        //2
	_        //skip the value 3
	o        //no need to specify iota
	_
	p
)

func main() {
	const i int = 12
	const j float32 = 32.1
	const k string = "fsffd"
	const l bool = true
	const m complex128 = complex(2, 5)

	// i = 13  cannot reassign values of constants

	fmt.Printf("%v, %T\n", i, i)
	fmt.Println(x)
	fmt.Println(y)
	fmt.Println(z)
	fmt.Println(o)
	fmt.Println(p)
}
