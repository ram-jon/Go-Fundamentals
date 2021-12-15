// Waiting using Wait Groups
package main

import (
	"fmt"
	"sync"
	"time"
)

// var wg = sync.WaitGroup{} //create  Wait Group

// func main() {
// 	message := "Initial Message"
// 	wg.Add(1) //increment counter
// 	go func() {
// 		fmt.Println(message)
// 		wg.Done() //decrement counter resumes block routines
// 	}()
// 	wg.Wait() //wait until counter is zero
// 	message = "change after calling routine"
// 	time.Sleep(100 * time.Millisecond)
// }

// -----------------------
// Example 2

func worker(id int, wg *sync.WaitGroup) {
	defer wg.Done() //defer this statement to the end of the function

	fmt.Printf("Worker %d starting\n", id)

	time.Sleep(time.Second)

	fmt.Printf("Worker %d done\n", id)
}

func main() {
	var wg sync.WaitGroup

	for i := 1; i <= 5; i++ {
		wg.Add(1)
		worker(i, &wg)
	}

	wg.Wait()
}
