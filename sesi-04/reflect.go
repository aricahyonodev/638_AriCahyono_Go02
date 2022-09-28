package main

import (
	"fmt"
	"reflect"
)

func main()  {
	var number = 23
	var reflectValue = reflect.ValueOf(number)

	fmt.Println("Tipe variabel :", reflectValue.Type())

	if reflectValue.Kind() == reflect.Int {
		fmt.Println("Nilai variable :", reflectValue.Int())
	}

	// Accessing value using interface method
	fmt.Println("")
	fmt.Println("Tipe variabel :", reflectValue.Type())
	fmt.Println("Nilai variabel :", reflectValue.Interface())

}