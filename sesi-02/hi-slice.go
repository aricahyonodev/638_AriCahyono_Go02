package main

import (
	"fmt"
	"strings"
)

func main()  {

	// Create Slice
	fmt.Println("Create Slice")
	var fruits = make([]string, 3)
	_ = fruits

	fruits[0] = "apple"
	fruits[1] = "banana"
	fruits[2] = "mango"
	fmt.Printf("%#v", fruits)

	// Create Slice With Append
	fmt.Println("\n\nCreate Slice With Append")
	var fruits2 = make([]string, 3)
	fruits2 = append(fruits2, "apple", "banana", "mango")
	fmt.Printf("%#v", fruits2)

	// Create Slice With Append  & Ellipsis 
	fmt.Println("\n\nCreate Slice With Append  & Ellipsis")
	var fruits3 = []string{"apple", "banana", "mango"}
	var fruits4 = []string{"durian", "pineapple", "starfruit"}
	fruits3 = append(fruits3, fruits4...)
	fmt.Printf("%#v", fruits3)

	// Slice Copy Function 
	fmt.Println("\n\nSlice Copy Function")
	var fruits5 = []string{"apple", "banana", "mango"}
	var fruits6 = []string{"durian", "pineapple", "starfruit"}
	nn := copy(fruits3, fruits4)

	fmt.Println("Fruits5 =>", fruits5)
	fmt.Println("Fruits6 =>", fruits6)
	fmt.Println("Copied elements =>", nn)

	// Slice Slicing 
	fmt.Println("\n\nSlice Slicing")
	var fruits7 = []string{"apple", "banana", "mango", "durian", "pineapple"}
	var slicingFruits1 = fruits7[1:4]
	fmt.Printf("%#v\n", slicingFruits1)
	var slicingFruits2 = fruits7[0:]
	fmt.Printf("%#v\n", slicingFruits2)
	var slicingFruits3 = fruits7[:3]
	fmt.Printf("%#v\n", slicingFruits3)
	var slicingFruits4 = fruits7[:]
	fmt.Printf("%#v\n", slicingFruits4)

	// Slice ( Combining & Append ) 
	fmt.Println("\n\nSlice ( Combining & Append )")
	var fruits8 = []string{"apple", "banana", "mango", "durian", "pineapple"}
	fruits8 = append(fruits8[:3], "rambutan")
	fmt.Printf("%#v\n", fruits8)

	// Slice ( Backing Array ) 
	fmt.Println("\n\nSlice ( Backing Array )")
	var fruits9 = []string{"apple", "banana", "mango", "durian", "pineapple"}
	var fruits10 = fruits9[2:4]
	fruits10[0] = "rambutan"
	fmt.Println("fruits 9 => ", fruits9)
	fmt.Println("fruits 10 => ", fruits10)

	// Slice ( Cap Array ) 
	fmt.Println("\n\nSlice ( Cap Array )")
	var fruits11 = []string{"apple", "banana", "mango", "banana"}
	fmt.Println("Fruits11 cap", cap(fruits11))
	fmt.Println("Fruits11 len", len(fruits11))
	fmt.Println(strings.Repeat("#", 20))

	var fruits12 = fruits11[0:3]
	fmt.Println("Fruits12 cap", cap(fruits12))
	fmt.Println("Fruits12 len", len(fruits12))
	fmt.Println(strings.Repeat("#", 20))

	var fruits13 = fruits11[1:]
	fmt.Println("Fruits13 cap", cap(fruits13))
	fmt.Println("Fruits13 len", len(fruits13))
	fmt.Println(strings.Repeat("#", 20))

	// Slice ( Creating a new backing ) 
	fmt.Println("\n\nSlice ( Creating a new backing ")
	cars := []string{"Ford", "Honda", "Audi", "Range Rover"}
	newCars := []string{}

	newCars =append(newCars, cars[0:2]...)
	cars[0] = "Nissan"
	fmt.Println("cars:", cars)
	fmt.Println("newCars:", newCars)

}