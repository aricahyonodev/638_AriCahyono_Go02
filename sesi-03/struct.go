package main

import (
	"fmt"
	"strings"
)

type Employee struct {
	name string
	age  int
	division string
} 

type Employee1 struct {
	name string
	age  int
	division string
} 

type Employee2 struct {
	name string
	age  int
	division string
} 

type Person struct {
	name string
	age int
}

type Employee3 struct {
	division string
	person Person
}

func main()  {

	// Struct (Giving value to struct)
	var employee Employee

	employee.name 	   = "Airell"
	employee.age  	   = 23
	employee.division  = "Curiculum Developer"
	fmt.Println(employee.name)
	fmt.Println(employee.age)
	fmt.Println(employee.division)

	// Struct (Initializing struct)
	var employee1 	  = Employee1{}
	employee1.name 	  = "Airell"
	employee.age 	  = 23
	employee.division = "Curriculum Developer"

	var employee2 = Employee1{name : "Ananda", age: 23, division: "Finance"}
	fmt.Printf("Employee1 : %+v\n", employee1)
	fmt.Printf("Employee2 : %+v\n", employee2)

	// Struct (Pointer To a Struct)
	var employee3 = Employee2{name : "Airell", age: 23, division: "Curiculum Developer"}
	var employee4 *Employee2 = &employee3

	fmt.Println("Employee3 name:", employee3.name)
	fmt.Println("Employee4 name:", employee4.name)

	fmt.Println(strings.Repeat("#", 20))

	employee3.name = "Ananda"

	fmt.Println("Employee3 name:", employee3.name)
	fmt.Println("Employee4 name:", employee4.name)

	// Struct (Embedded struct)
	var employee5 		  = Employee3{}
	employee5.person.name = "Airell"
	employee5.person.age  = 23
	employee.division     = "Curriculum Developer"
	fmt.Printf("%+v", employee5)

	// Struct (Anonymous struct)
	var employee6 = struct {
		person Person
		division string
	}{}

	employee6.person   = Person{name: "Airell", age : 23}
	employee6.division = "Curriculum Developer"

	// Anonymous struct dengan perngisian field
	var employee7 = struct {
		person Person
		divison string
	}{
		person: Person{name: "Ananda", age : 23},
		divison: "Finance",
	}

	fmt.Printf("\nEmployee6: %+v\n", employee6)
	fmt.Printf("Employee7: %+v\n", employee7)


	// Struct (Slice Of Struct)
	var people = []Person{
		{name: "Airell", age: 23},
		{name: "Ananda", age: 23},
		{name: "Mailo", age: 23},
	}

	for _, v := range people {
		fmt.Printf("%+v\n", v)
	}


	// Struct (Slice of Anonymous struct)
	var employee8 = []struct {
		person Person
		division string
	}{
		{person: Person{name: "Airell", age: 23}, division: "Curriculum Developer"},
		{person: Person{name: "Ananda", age: 23}, division: "Finance"},
		{person: Person{name: "Ananda", age: 23}, division: "Marketing"},
	}

	for _, em := range employee8 {
		fmt.Printf("%+v\n", em)
	}

}