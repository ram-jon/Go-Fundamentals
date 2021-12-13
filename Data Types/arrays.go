package main

import "fmt"

func main() {
	var arr [3]int = [3]int{1, 2, 3}
	amt := [3]int{23, 56, 324}                             //shorthand for declaring array
	salary := [...]int{32423, 65342, 234554, 654323, 7643} //specify "..." insted of size

	fmt.Println(arr) //print array
	fmt.Println(amt)
	fmt.Println(salary)
	// print length of array using len function
	fmt.Println(len(salary))

	// set and get value of perticular index
	arr[0] = 20
	fmt.Println(arr[0])

	// declare array and assign values later
	var arr1 [4]string
	arr1[0] = "qwe"
	arr1[1] = "adad"
	arr1[2] = "dasf"

	//Iterate over array using loops
	for x := 0; x < len(arr1); x++ {
		fmt.Printf("%s\n", arr1[x])
	}

	// copy one array into another
	strArray := arr1

	fmt.Println(strArray)

	//changes made in new array variable does not affect original array
	//array variable is not reference type in go
	strArray[0] = "abc"

	fmt.Printf("Original: %v \n New : %v\n", arr1, strArray)

	// copy reference of the array into pointer variable
	ptrVar := &strArray //it will create pointer variable ptrVar and assign address of strArray

	ptrVar[0] = "new val" //it will change both ptrVar and strArray

	fmt.Printf("ptrVar: %v \n strArray : %v\n", ptrVar, strArray)

	fmt.Printf("%T", ptrVar)

	// Slicing of Array
	fmt.Println("--------------------------------------------------------------")
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

	// Multi dimentional array
	var matrix [3][3]int = [3][3]int{
		[3]int{1, 0, 0},
		[3]int{0, 1, 0},
		[3]int{0, 0, 1},
	}
	matrix[0][1] = 2
	fmt.Println(matrix)

	//sorthand declaration
	matrix1 := [2][2]int{{1, 2},
		{3, 4}}
	fmt.Println(matrix1)
}
