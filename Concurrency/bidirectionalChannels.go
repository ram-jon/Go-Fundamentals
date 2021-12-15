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
		ch <- 10 //send value to the channel
		wg.Done()
	}()

	go func() {
		ch <- 23          //send value to the channel
		fmt.Println(<-ch) //get value from channel
		wg.Done()
	}()

	wg.Wait()
}
