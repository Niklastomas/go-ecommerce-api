package api

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/niklastomas/go-ecommerce-api/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type Server struct {
	DB     *gorm.DB
	Router *mux.Router
}

func (server *Server) Init() (err error) {
	dsn := "host=localhost user=postgres password=postgres dbname=ecommerceDB port=5432 sslmode=disable"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Println(err)
		return
	}

	server.DB = db
	server.Router = mux.NewRouter()

	err = server.DB.AutoMigrate(&models.User{}, &models.Product{}, &models.Category{}, &models.Order{}, &models.OrderItem{}, &models.Payment{})
	if err != nil {
		log.Println(err)
		return
	}

	return nil

}

func (server *Server) Run(port string) error {
	fmt.Printf("Server is running on port %s", port)
	err := http.ListenAndServe(port, server.Router)
	if err != nil {
		return err
	}
	return nil
}
