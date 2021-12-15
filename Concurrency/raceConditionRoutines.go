// A race condition occurs when multiple threads try to access and modify the same data (memory address). E.g., if one thread tries to increase an integer and another thread tries to read it, this will cause a race condition.

package main

import (
	"fmt"
	"time"
)

//Detect data race in program using "-race" flag while running program
// func main() {
// 	message := "Initial Message"
// 	go func() {
// 		fmt.Println(message)   //this will create race condition
// 	}()
// 	message = "change after calling routine"
// 	time.Sleep(100 * time.Millisecond)
// }

//---------------------------------------

func main() {
	message := "Initial Message"
	go func(message string) {
		fmt.Println(message) //this will not create race condition beacause func reading parameter value passed to it
	}(message)
	message = "change after calling routine"
	time.Sleep(100 * time.Millisecond)
}
