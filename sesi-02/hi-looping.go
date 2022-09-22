package main

import "fmt"

func main()  {

	// First Way of Looping
	fmt.Println("First Way of Looping")
	for angka := 0; angka < 3; angka++ {
		fmt.Println("Angka", angka)
	}

	// Second Way of Looping
	fmt.Println("Second Way of Looping")
	var i = 0
	for i < 3  {
		fmt.Println("Angka", i)
		i++
	}

	// Third Way of Looping
	fmt.Println("Third Way of Looping")
	var u = 0
	for {
		fmt.Println("Angka", u)
		u++
		if u == 3 {
			break
		}
	}

	// Break & Continue
	fmt.Println("Break & Continue")
	for e := 1; e <= 10; e++ {
		if (e % 2) == 1 {
			continue
		}

		if 3 > 8 {
			break
		}

		fmt.Println("Angka ", e)
	}

	// Nested Looping
	fmt.Println("Nested Looping")
	for a := 0; a < 5; a++ {
		for b := a; b < 5; b++ {
			fmt.Print(b, " ")
		}
			fmt.Println()
	}

	// Looping Label
	fmt.Println("Looping Label")
	outerLoop:
	for c := 0; c < 3; c++ {
		fmt.Println("Perulangan ke - ", c+1)
		for d := 0; d < 3; d++ {
			if i == 2{
				break outerLoop
			}
			fmt.Print(d, " ")
		}
		fmt.Println("")
	}
}