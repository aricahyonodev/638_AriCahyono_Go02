package main

import (
	"fmt"
	"strings"
)

func main()  {

	// Learn Array
	fmt.Println("Learn Array")
	var number [4]int
	number = [4]int{1,2,3,4}
	var stringsArray = [4]string{"Airell", "Nanda", "Mailo"}

	fmt.Printf("%v\n", number)
	fmt.Printf("%v\n", stringsArray)

	// Array Modify
	fmt.Println("\nArray Modify")
	var fruits = [3]string{"apel", "pisang", "mangga"}
	fruits[0] = "apple"
	fruits[1] = "banana"
	fruits[2] = "mango"
	fmt.Printf("%v\n", fruits)

	// Array Looping
	fmt.Println("\nArray Looping")
	var fruits2 = [3]string{"apel", "pisang", "mangga"}
	fmt.Println("\nFirst Way")
	for i, v := range fruits2 {
		fmt.Printf("Index: %d, Values: %s\n", i, v)
	}

	fmt.Println(strings.Repeat("#", 25))
	fmt.Println("Second Way")
	for i := 0; i < len(fruits2); i++ {
		fmt.Printf("Index: %d, Values: %s\n", i, fruits2[i])
	}

	// Nested Array
	fmt.Println("\nNested Array")
	balances := [2][4]int{{5,6,7,8},{9,10,11,12}}

	for _, arr := range balances{
		fmt.Printf("%d \n", arr)
		for _, value := range arr {
			fmt.Printf("%d \n ", value)
		}
	}
}