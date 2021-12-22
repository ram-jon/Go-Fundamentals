package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	// using fmt
	// it will not take space separated values
	// var name string
	// fmt.Scanln(&name)
	// fmt.Println(name)

	// reader := bufio.NewReader(os.Stdin)
	// fmt.Println("Enter Full Name")
	// mystring, _ := reader.ReadString('\n')
	// fmt.Println(mystring)

	// taking number input

	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Enter Age")
	age, _ := reader.ReadString('\n') //take input in string format
	//trim spaces from string and convert into int64 with base 10
	ageNum, _ := strconv.ParseInt(strings.TrimSpace(age), 10, 64)
	fmt.Println(ageNum + 2)

}
