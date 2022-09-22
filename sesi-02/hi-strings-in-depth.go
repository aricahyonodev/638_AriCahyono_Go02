package main

import (
	"fmt"
	"unicode/utf8"
)

func main()  {

	// Looping Over String
	fmt.Println("\n\nLooping Over String")
	city := "Jakarta"
	for i := 0; i < len(city); i++ {
		fmt.Printf("index: %d, byte: %d\n", i, city[i])
	}

	// Convert Byte to String
	fmt.Println("\n\nConvert Byte to String")
	var city2 []byte = []byte{74, 97, 107, 97, 114, 116, 97}
	fmt.Println(string(city2))

	// Compare String to Byte
	fmt.Println("\n\nCompare String to Byte")
	str1 := "makan"
	str2 := "mânca"
	fmt.Printf("str1 byte length => %d\n", len(str1))
	fmt.Printf("str2 byte length => %d\n", len(str2))

	// Compare String to Byte to String
	fmt.Println("\n\nCompare String to Byte to String")
	fmt.Printf("str1 character length => %d\n", utf8.RuneCountInString(str1))
	fmt.Printf("str2 character length => %d\n", utf8.RuneCountInString(str2))

	// Looping String
	fmt.Println("\n\nLooping String")
	for i, s := range str2 {
		fmt.Printf("index => %d, string => %s\n", i, string(s))
	}
	
}