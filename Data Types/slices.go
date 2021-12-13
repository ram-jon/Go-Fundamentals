package main

import "fmt"

func main() {
	var x []int = []int{1, 2, 3}
	y := x

	// slices arr reference type variale it is pointing to underlying array
	y[0] = 3 //it will reflect in both a and b
	fmt.Println(x)
	fmt.Println(y)

	//sorthand declaration
	mySlice := []int{4, 5, 6, 3}
	fmt.Println(mySlice)

	// length of slice
	fmt.Println(len(mySlice))

	//capacity of slice
	fmt.Println(mySlice)

	// create slices using array
	a := [...]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	b := a[:]   //assign element from index 0 to end to b
	c := a[2:]  //start from index 2 to the end index of array
	d := a[:5]  //start from index 0 upto index 4 exclude index 5
	e := a[2:7] //start from 2 to index 6 exclude ending index 7
	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(c)
	fmt.Println(d)
	fmt.Println(e)

	// Create slice using make function
	// Syntax : make(type, size, capacity)

	var newSlice []int = make([]int, 3, 6)
	fmt.Printf("Size : %v , Capacity : %v\n", len(newSlice), cap(newSlice))

	// Append data to slice

	var slice1 []int = []int{1, 2, 3}
	var slice2 []int = append(slice1, 4)
	var slice3 []int = append(slice1[1:], 5)
	var slice4 []int = append(slice2, slice3...) //append all elemnts of slice3 to slice2 using spread operator
	fmt.Println(slice1)
	fmt.Println(slice2)
	fmt.Println(slice3)
	fmt.Println(slice4)

}
