package main

import (
	"fmt"
	"strings"
	"math"
)

func main(){

	greet("Airell", 23)


	// Function Return
	var names = []string{"Airell", "Jordan"}
	var printMsg = greet2("Heii", names)
	fmt.Println("")
	fmt.Println(printMsg)

	// Function Returning Multiple Values
	var diameter float64 = 15
	var area, circumFerence = calculate(diameter)
	fmt.Println("\nArea:", area)
	fmt.Println("Circumference", circumFerence)

	// Variadic Function 1
	fmt.Println("")
	studentLists := print("Airell", "Nanda", "Mailo", "Schannel", "Marco")
	fmt.Printf("%v", studentLists)

	// Variadic Function 2
	numberLists := []int{1,2,3,4,5,6,7,8}
	result := sum(numberLists...)
	fmt.Println("")
	fmt.Println("Result:", result)

	// Variadic Function 3
	fmt.Println("")
	profile("Airell", "pasta", "ayam geprek", "ikan roa", "sate padang")
}

func greet(name string, age int8){
	fmt.Printf("Hello there! my name is %s and I'm %d years old", name, age)
}
func greet2(msg string, names []string) string{
	var joinStr = strings.Join(names, "")
	var result string = fmt.Sprintf("%s %s", msg, joinStr)
	return result
}

func print(names ...string) []map[string]string {
		var result []map[string]string

		for i, v := range names {
			key := fmt.Sprintf("student %d", i+1)
			temp := map[string]string{
				key : v,
			}
			result = append(result, temp)
		}
		return result
}

func sum(numbers ...int) int{
	total := 0

	for _, v := range numbers {
		total += v
	}
	return total
}

func profile(name string, favFoods ...string){
	mergeFavFoods := strings.Join(favFoods,",")

	fmt.Println("Hello there!!! I'm", name)
	fmt.Println("I really love to eat", mergeFavFoods)
}

func calculate(d float64) (float64, float64){
	var area float64 = math.Pi * math.Pow(d/2, 2)

	var circumFerence = math.Pi * d

	return area, circumFerence
}

