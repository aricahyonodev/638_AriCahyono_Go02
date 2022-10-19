package main

import (
	"my-gram/routers"
)

func main() {
	r := router.StartApp()
	r.Run(":3000")
}