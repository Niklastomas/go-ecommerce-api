package main

import (
	"github.com/niklastomas/go-ecommerce-api/api"
)

func main() {
	server := api.Server{}

	if err := server.Init(); err != nil {
		panic(err)
	}

	server.InitRoutes()

	if err := server.Run(":8080"); err != nil {
		panic(err)
	}

}
