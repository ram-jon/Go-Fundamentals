package main

import (
	"fmt"
	"sync"
)

var wg = sync.WaitGroup{}

func main() {
	ch := make(chan int, 100) //creating buffered channel with buffer size 100
	wg.Add(2)

	// sending multiple elemnts and receiving only one

	// go func(ch <-chan int) {
	// 	i := <-ch //receiving only one item if its non buffered channel it needs to receive all items
	// 	fmt.Println(i)
	// 	wg.Done()
	// }(ch)

	// go func(ch chan<- int) {
	// 	// sending multiple data
	// 	ch <- 23
	// 	ch <- 34
	// 	ch <- 50
	// 	wg.Done()
	// }(ch)

	// sending and recieving multiple elements
	// go func(ch <-chan int) {
	// 	i := <-ch //receiving only one item if its non buffered channel it needs to receive all items
	// 	fmt.Println(i)
	// 	i = <-ch
	// 	fmt.Println(i)
	// 	wg.Done()
	// }(ch)

	// go func(ch chan<- int) {
	// 	// sending multiple data
	// 	ch <- 23
	// 	ch <- 34
	// 	wg.Done()
	// }(ch)

	// using for loop to send and receive channel data

	go func(ch <-chan int) {
		// receiving element using for range operator on channels

		// for i := range ch {
		// 	fmt.Println(i)
		// }

		// cheking the status of channel close or open using ",ok"

		for {
			if i, ok := <-ch; ok {
				fmt.Println(i)
			} else {
				break
			}
		}

		wg.Done()
	}(ch)

	go func(ch chan<- int) {
		// sending multiple data
		ch <- 23
		ch <- 34
		close(ch) //channel must be close to inform for loop where to stop
		wg.Done()
	}(ch)

	wg.Wait()
}
