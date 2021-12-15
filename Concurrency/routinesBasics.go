package main

import (
	"fmt"
	"time"
)

// func main() {
// 	go writeMessage()   //this will not print any message, the main function does not wait to complete execution of go routine
// }

func main() {
	go writeMessage()
	time.Sleep(100 * time.Millisecond) //we pause main method to complete the execution of routine

	// Anonymous go routines
	go func() {
		fmt.Println("Hello World")
	}()
	time.Sleep(100 * time.Millisecond)
}

func writeMessage() {
	fmt.Println("Inside routine")
}
