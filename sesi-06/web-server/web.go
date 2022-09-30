package main

import (
	"fmt"
	"net/http"
)

var PORT = ":8000"

func main()  {
	http.HandleFunc("/", greet)
	http.ListenAndServe(PORT, nil)
}

func greet(w http.ResponseWriter, r *http.Request){
	msg := "Hello World"
	fmt.Fprint(w, msg)
}

