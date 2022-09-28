package main

import (
	"fmt"
	"sync"
)

func main()  {
	fruits := []string{"apple", "manggo", "durian", "rambutan"}

	var wg sync.WaitGroup

	for i, f := range fruits {
		wg.Add(1)
		go printFruit(i, f, &wg)
	}

	wg.Wait()
}

func printFruit(index int, fruit string, wg *sync.WaitGroup){
	fmt.Printf("index => %d, fruit => %s \n", index, fruit)
	wg.Done()
}