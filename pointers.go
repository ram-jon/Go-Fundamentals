package main

import "fmt"

type Foo struct {
	bar int
}

func main() {
	var a int = 12
	var b *int = &a //create pointer variable b and assign address of a to it
	fmt.Println(a)
	fmt.Println(b) //prints the address of b
	//to print value of pointer we need to dereference it using "*" operator
	fmt.Println(*b) //print value of b

	// Pointers with structs
	var foo *Foo     //create pointer of type struct Foo
	fmt.Println(foo) //it will print <nil> default value
	foo = new(Foo)
	fmt.Println(foo)  //print address of struct initialize with 0
	fmt.Println(*foo) //print value

	// set and get properties of struct using pointer
	(*foo).bar = 35
	fmt.Println((*foo).bar)

	// Another method wihout "*"
	foo.bar = 55
	fmt.Println(foo.bar)

}
