// In Go language, a channel is a medium through which a goroutine communicates with another goroutine

package main

import (
	"fmt"
	"sync"
)

var wg = sync.WaitGroup{}

func main() {
	ch := make(chan int) //create channel of type int
	wg.Add(2)            //add counter 2 for 2 routines

	go func() {
		i := <-ch //get value from channel
		fmt.Println(i)
		wg.Done()
	}()

	go func() {
		ch <- 23 //send value to the channel
		wg.Done()
	}()

	wg.Wait()
}
