package main

import (
	"log"

	"github.com/joho/godotenv"
	"github.com/niklastomas/go-ecommerce-api/api"
)

func init() {
	err := godotenv.Load()
	if err != nil {
		log.Println(err)
	}
}

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
