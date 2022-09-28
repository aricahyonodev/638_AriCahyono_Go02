package main

import (
	"errors"
	"fmt"
	"github.com/howeyc/gopass"
)

func main()  {

// 1. buat login
// 2. require username & password
// 3. username > string
// 4. password > **** (len)

// if username && password  == true
// username dan password benar
// else
// username/password salah

	var username string
	var password []byte

	defer cathErr()

	fmt.Print("Masukkan Username : ")
	fmt.Scanln(&username)
	fmt.Scanln(username)
	fmt.Print("Masukkan Password : ")
	password, _ = gopass.GetPasswdMasked()

	 if valid, err := validUsernamePassword(username ,string(password)); err != nil {
		panic(err.Error())
	 }else{
		fmt.Println(valid)
	 }

}


func validUsernamePassword(username string,password string)(string, error){
	ul := len(username)
	pl := len(password)
	if pl < 5 && ul < 5{
		return "", errors.New("username/password salah")
	}else if pl < 5 {
		return "", errors.New("password has to have more than 4 charachter")
	}else if ul < 5 {
		return "", errors.New("username has to have more than 4 charachter")
	}
	return "username dan password benar", nil
}

func cathErr(){
	if r := recover(); r != nil {
		fmt.Println("Error occured:", r)
	}else{
		fmt.Println("Application running pervectly")
	}
}