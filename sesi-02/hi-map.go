package main

import "fmt"

func main()  {

	// Introduce Map
	fmt.Println("\n\nIntroduce Map")
	var person map[string]string // Deklarasi
	person = map[string]string{} // Inisialisasi

	person["name"] 	   = "Ari"
	person["age"]  	   = "24"
	person["address"]  = "Jalan Sudirman"

	fmt.Println("Name: ", person["name"])
	fmt.Println("Age: ", person["age"])
	fmt.Println("Address: ", person["address"])


	// Map With Value
	fmt.Println("\n\nMap With Value")
	var person2 = map[string]string{
		"name" : "ari",
		"age" : "24",
		"address" : "Jalan Sudirman",
	}
	fmt.Println("Name: ", person2["name"])
	fmt.Println("Age: ", person2["age"])
	fmt.Println("Address: ", person2["address"])

	// Looping Map
	fmt.Println("\n\nLooping Map")
	for k, v := range person2 {
		fmt.Println(k, ":", v)
	}

	// Delete Map
	fmt.Println("\n\nDelete Map")
	fmt.Println("Before Deleting", person2)
	delete(person2, "age")
	fmt.Println("After Deleting", person2)



	// Map ( Detectng a value )
	fmt.Println("\n\nMap ( Detectng a value )")
	var person3 = map[string]string{
		"name" : "ari",
		"age" : "24",
		"address" : "Jalan Sudirman",
	}

	value,exist := person3["age"]

	if exist {
		fmt.Println(value)
	}else{
		fmt.Println("vaue does'nt exist")
	}

	// Map ( Combining slice )
	fmt.Println("\n\nMap ( Combining slice )")
	var people4 = []map[string]string{
		{"name": "Airell", "age" : "23" },
		{"name": "Nanda", "age" : "23" },
		{"name": "Mailo", "age" : "15" },
	}
	for i, person4 := range people4 {
		fmt.Printf("Index : %d, name: %s, age: %s\n", i, person4["name"], person4["age"])
	}

}