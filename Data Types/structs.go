package main

import "fmt"

type Student struct {
	name     string
	rollNo   int
	subjects []string
}

func main() {
	// Declaring a variable of a `struct` type
	// All the struct fields are initialized
	// with their zero value
	var a Student
	fmt.Println(a)
	student1 := Student{
		name:   "Ram",
		rollNo: 12,
		subjects: []string{
			"Maths",
			"Physics",
			"Chemistry",
		},
	}

	//Shorthand declaration not preffered
	//you need to specify value for every property of struct in order
	student2 := Student{
		"Sham",
		4,
		[]string{
			"Biology",
			"Electronics",
		},
	}

	fmt.Println(student1)
	fmt.Println(student2)

	// Get and Set single property from struct

	student1.rollNo = 1          //set rollno
	fmt.Println(student1.rollNo) //get rollno from struct
	fmt.Println(student1.subjects[0])

	// copying structs
	// Note : struct are value type not reference type like slices

	student3 := student1
	// student3 := &student //it will copy address (reference)
	student3.name = "Amit" //it will reflect only in student3 not in student1

	fmt.Println(student1)
	fmt.Println(student3)
}
