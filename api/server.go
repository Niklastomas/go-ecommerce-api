package api

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/gorilla/mux"
	"github.com/niklastomas/go-ecommerce-api/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type Server struct {
	DB              *gorm.DB
	Router          *mux.Router
	ProductHandler  ProductHandler
	UserHandler     UserHandler
	CategoryHandler CategoryHandler
	OrderHandler    OrderHandler
	PaymentHandler  PaymentHandler
	AuthHandler     AuthHandler
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

	// migrate
	err = server.DB.AutoMigrate(
		&models.User{},
		&models.Product{},
		&models.Category{},
		&models.Order{},
		&models.OrderItem{},
		&models.Payment{},
	)
	if err != nil {
		log.Println(err)
		return
	}

	// logger
	l := log.New(os.Stdout, "e-commerce-api", log.LstdFlags)

	// create handlers
	server.ProductHandler = NewProducts(l, server.DB)
	server.UserHandler = NewUsers(l, server.DB)
	server.CategoryHandler = NewCategories(l, server.DB)
	server.OrderHandler = NewOrders(l, db)
	server.PaymentHandler = NewPayments(l, db)
	server.AuthHandler = NewAuth(l, db)

	return nil

}

func (server *Server) Run(addr string) error {
	s := &http.Server{
		Addr:         addr,
		Handler:      server.Router,
		ReadTimeout:  time.Second * 10,
		WriteTimeout: time.Second * 10,
		IdleTimeout:  time.Second * 15,
	}

	fmt.Printf("Server is running on %s", addr)
	err := s.ListenAndServe()
	if err != nil {
		return err
	}
	return nil
}
