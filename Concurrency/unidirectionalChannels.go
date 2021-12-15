package main

import (
	"fmt"
	"sync"
)

var wg = sync.WaitGroup{}

func main() {
	ch := make(chan int) //create bidirectional channel of type int

	//we can also create unidirectional channel
	// Only for receiving
	// mychanl1 := make(<-chan string)

	// Only for sending
	// mychanl2 := make(chan<- string)

	wg.Add(2) //add counter 2 for 2 routines

	go func(ch <-chan int) { //channel is converted to receive only
		i := <-ch //get value from channel
		fmt.Println(i)
		// ch <- 10   //it will give error channel receive only
		wg.Done()
	}(ch)

	go func(ch chan<- int) { //channel pass as send only
		ch <- 23 //send value to the channel
		// fmt.Println(<-ch) //it will give error cause channel is send only
		wg.Done()
	}(ch)

	wg.Wait()
}
