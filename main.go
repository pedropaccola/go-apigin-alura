package main

import (
	"github.com/pedropaccola/go-apigin-alura/database"
	"github.com/pedropaccola/go-apigin-alura/routes"
)

func main() {
	database.ConnectDB()
	routes.HandleRequests()
}
