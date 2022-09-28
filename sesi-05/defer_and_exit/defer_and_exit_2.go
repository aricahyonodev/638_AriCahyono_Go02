package main

import (
	"fmt"
)

func main()  {
	callDeferFunc()
	fmt.Println("Welcome back to Go learning center")
}

func callDeferFunc(){
	defer deferFunc()
}

func deferFunc(){
	fmt.Println("Defer func starts to execute")
}