package main

import "fmt"

func main()  {

	// Pointer Memori Address
	var firstNumber int = 4
	var secondNumber *int = &firstNumber
	fmt.Println("firstNumber (Value) : ", firstNumber)
	fmt.Println("firstNumber (Memori address) : ", &firstNumber)

	fmt.Println("secondNumber (Value) : ", *secondNumber)
	fmt.Println("secondNumber (Memori address) : ", secondNumber)

	// PointerChanging value through pointer
	var firstPerson string   = "John"
	var secondPerson *string = &firstPerson

	fmt.Println("firstPerson (value) : ", firstPerson)
	fmt.Println("firstPerson (memori address) : ", &firstPerson)
	fmt.Println("secondPerson (value) : ", *secondPerson)
	fmt.Println("secondPerson (memori address) : ", secondPerson)

	*secondPerson = "Doe"

	fmt.Println("firstPerson (value) : ", firstPerson)
	fmt.Println("firstPerson (memori address) : ", &firstPerson)
	fmt.Println("secondPerson (value) : ", *secondPerson)
	fmt.Println("secondPerson (memori address) : ", secondPerson)

	var a int = 10
	fmt.Println("Before:", a)
	changeValue(&a)
	fmt.Println("After:", a)
	
}

func changeValue(number *int){
	*number = 20
}