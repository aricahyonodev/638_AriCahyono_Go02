package main

import "sesi-03/helpers"

func main()  {

	helpers.Greet()

	var person = helpers.Person{}
	person.InvokeGreet()
}