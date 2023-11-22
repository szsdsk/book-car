package main

import (
	"acs/src/database"
	"acs/src/routes"
)

func main() {
	database.Connect()
	routes.InitRouter()
}
