package main

import (
	"fmt"
)

func main()  {

	// Learn If/Else Condition
	fmt.Println("Learn If/Else Condition")
	var currentYear = 2000
	if age := currentYear - 1998; age < 17  {
		fmt.Println("Kamu belum boleh membuat sim")
	}else{
		fmt.Println("Kamu sudah boleh membuat sim")
	}

	// Learn Switch/Case Condition
	fmt.Println("\nLearn Switch/Case Condition")
	var score = 8
	switch score {
	case 8:
		fmt.Println("perfect")
	case 7:
		fmt.Println("awesome")
	default:
		fmt.Println("not bad")
	}

	// Learn fallthrough Switch/Case Condition
	fmt.Println("\nLearn fallthrough Switch/Case Condition")
	var score2 = 4
	switch {
	case score2 == 8:
		fmt.Println("perfect")
	case (score2 < 8) && (score2 > 3):
		fmt.Println("not bad")
		fallthrough
	case score2 < 5:
		fmt.Println("It is ok, but please study harder")
	default:
		{
			fmt.Println("Study harder")
			fmt.Println("You don't have a good score yet")
		}
	}

	// Learn Nested Condition
	fmt.Println("\nLearn fallthrough Switch/Case Condition")
	var score3 = 10

	if score3 > 7 {
		switch score3{
		case 10: 
			fmt.Print("perfect")
		default:
			fmt.Println("nice!")		
		}
	}else{
		if score3 == 5 {
			fmt.Println("not bad")
		}else if score3 == 3 {
			fmt.Println("keep trying")
		}else{
			fmt.Println("you can do it")
			if score3 == 0 {
				fmt.Println("try harder!")
			}
		}
	}
}

