package main

import (
	"errors"
	"fmt"
	"strconv"
)

func main()  {
	var number int
	var err error

	number, err = strconv.Atoi("123GH")
	 if err == nil {
		fmt.Println(number)
	 }else{
		fmt.Println(err.Error())
	 }

	 number, err = strconv.Atoi("123")
	 if err == nil {
		fmt.Println(number)
	 }else{
		fmt.Println(err.Error())
	 }

	 defer cathErr()
	 var password string
	 fmt.Scanln(&password)

	 if valid, err := validPassword(password); err != nil {
		panic(err.Error())
	 }else{
		fmt.Println(valid)
	 }

}

func validPassword(password string)(string, error){
	pl := len(password)
	if pl < 5 {
		return "", errors.New("password has to have more than 4 charachter")
	}
	return "Valid password", nil
}

func cathErr(){
	if r := recover(); r != nil {
		fmt.Println("Error occured:", r)
	}else{
		fmt.Println("Application running pervectly")
	}
}