package main

import (
	"fmt"
	"time"
)

func main()  {

	// buffered channel.
	c1 := make(chan int, 3)

	go func(c chan int){
		for i := 1; i <= 5; i++ {
			fmt.Println("func goroutines starts sending data into the channel")
			c <- i
			fmt.Println("func goroutine after sending data into the channel")
		}
		close(c)
	}(c1)

	fmt.Println("Main goroutine sleeps for 2 seconds")
	time.Sleep(time.Second * 2)

	for v := range c1 {
		fmt.Println("main goroutine received value from channel:", v)
	}
}
