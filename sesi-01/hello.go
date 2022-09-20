package main

import "fmt"

func main()  {

	// Go Introduction
	fmt.Print("Hello GO, lets start my journey \n")

	// Variable
	var name    = "Ari Cahyono"
	var age int = 24
	city 		:= "Kota Surabaya"
	fmt.Printf("Nama saya adalah %s, Umur %d tahun, dari %s \n", name, age, city)

	var variableNotUse = ""
	_ = variableNotUse; 

	first, second 	  := 1,2
	var third, fourth = 3,4
	fmt.Println(first, second)
	fmt.Println(third, fourth)


	// Const & Operator
	const full_name = "Ari Cahyono"
	fmt.Printf("Hello %s \n", full_name)

	var value = 2 + 3
	fmt.Println(value)

	var firstCondition bool  = 2 < 3
	var secondCondition bool = "joey" == "Joey"
	var thirdCondition bool  = 10 != 2.3
	var fourthCondition bool = 11 <= 11

	fmt.Println("First Condition : ", firstCondition)
	fmt.Println("Second Condition : ", secondCondition)
	fmt.Println("Third Condition : ", thirdCondition)
	fmt.Println("Fourth Condition : ", fourthCondition)


	var right = true
	var wrong = false
	
	var wrongAndRight = right && wrong
	fmt.Printf("Wrong && Right \t(%t)\n", wrongAndRight)

	var wrongOrRight = right || wrong
	fmt.Printf("Wrong OR Right \t(%t)\n", wrongOrRight)

	var wrongReverse = !wrong
	fmt.Printf("!left\t(%t)\n", wrongReverse)
}