package main

import (
	"go_jwt/database"
	"go_jwt/router"
)

func main() {
	database.StartDB()

	r 	:= router.StartApp()
	r.Run(":3000")

}