package main

import (
	"ass3_be/router"
)

func main() {
	r := router.StartApp()
	r.Run(":3000")
}