package api

import (
	"fmt"
	"net/http"

	"github.com/niklastomas/go-ecommerce-api/middleware"
)

func (s *Server) InitRoutes() {

	s.Router.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "Hello")
	}).Methods("GET")

	// users
	s.Router.HandleFunc("/api/users", middleware.JwtMiddleware(s.UsersLIST)).Methods("GET")
	s.Router.HandleFunc("/api/users/{id}", middleware.JwtMiddleware(s.GetUserById)).Methods("GET")
	s.Router.HandleFunc("/api/users/{id}", middleware.JwtMiddleware(s.UpdateUser)).Methods("PUT")

	// auth
	s.Router.HandleFunc("/api/login", s.Login).Methods("POST")
	s.Router.HandleFunc("/api/register", s.Register).Methods("POST")

	//products
	s.Router.HandleFunc("/api/products", s.CreateProduct).Methods("POST")
	s.Router.HandleFunc("/api/products", s.GetAllProducts).Methods("GET")
	s.Router.HandleFunc("/api/products/{id}", s.GetProductById).Methods("GET")
	s.Router.HandleFunc("/api/products/{id}", s.UpdateProduct).Methods("PUT")

	// category
	s.Router.HandleFunc("/api/category", s.GetAllCategories).Methods("GET")
	s.Router.HandleFunc("/api/category", s.CreateCategory).Methods("POST")
	s.Router.HandleFunc("/api/category/{id}", s.GetCategoryById).Methods("GET")
	s.Router.HandleFunc("/api/category/{id}", s.DeleteCatecory).Methods("DELETE")
	s.Router.HandleFunc("/api/category/{id}", s.UpdateCategory).Methods("PUT")

	// orders
	s.Router.HandleFunc("/api/orders", s.CreateOrder).Methods("POST")
	s.Router.HandleFunc("/api/orders", s.GetAllOrders).Methods("GET")
	s.Router.HandleFunc("/api/orders/{id}", s.UpdateOrder).Methods("PUT")
	s.Router.HandleFunc("/api/orders/{id}", s.GetOrderById).Methods("GET")

	// payment
	s.Router.HandleFunc("/api/charge/{orderId}", s.Charge).Methods("POST")

}
