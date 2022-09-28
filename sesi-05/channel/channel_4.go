package main

import (
	"fmt"
	"time"
)

func main()  {

	// Unbuffered channel.
	c1 := make(chan int)

	go func(c chan int){
		fmt.Println("func goroutines starts sending data into the channel")
		c <- 10
		fmt.Println("func goroutine after sending data into the channel")
	}(c1)

	fmt.Println("Main goroutine sleeps for 2 seconds")
	time.Sleep(time.Second * 2)

	fmt.Println("Main goroutine starts receivng data")
	d := <- c1
	fmt.Println("main goroutine starts receiving data", d)

	close(c1)
	time.Sleep(time.Second)
}
