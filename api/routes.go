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
	s.Router.HandleFunc("/api/users", middleware.Jwt(s.UserHandler.GetUsers)).Methods("GET")
	s.Router.HandleFunc("/api/users/{id}", middleware.Jwt(s.UserHandler.GetUserById)).Methods("GET")
	s.Router.HandleFunc("/api/users/{id}", middleware.Jwt(s.UserHandler.UpdateUser)).Methods("PUT")

	// auth
	s.Router.HandleFunc("/api/login", s.AuthHandler.Login).Methods("POST")
	s.Router.HandleFunc("/api/register", s.AuthHandler.Register).Methods("POST")

	//products
	s.Router.HandleFunc("/api/products", middleware.Cors(s.ProductHandler.CreateProduct)).Methods("POST")
	s.Router.HandleFunc("/api/products", middleware.Cors(s.ProductHandler.GetAllProducts)).Methods("GET")
	s.Router.HandleFunc("/api/products/{id}", middleware.Cors(s.ProductHandler.GetProductById)).Methods("GET")
	s.Router.HandleFunc("/api/products/{id}", middleware.Cors(s.ProductHandler.UpdateProduct)).Methods("PUT")

	// category
	s.Router.HandleFunc("/api/category", s.CategoryHandler.GetAllCategories).Methods("GET")
	s.Router.HandleFunc("/api/category", s.CategoryHandler.CreateCategory).Methods("POST")
	s.Router.HandleFunc("/api/category/{id}", s.CategoryHandler.GetCategoryById).Methods("GET")
	s.Router.HandleFunc("/api/category/{id}", s.CategoryHandler.DeleteCatecory).Methods("DELETE")
	s.Router.HandleFunc("/api/category/{id}", s.CategoryHandler.UpdateCategory).Methods("PUT")

	// orders
	s.Router.HandleFunc("/api/orders", s.OrderHandler.CreateOrder).Methods("POST")
	s.Router.HandleFunc("/api/orders", s.OrderHandler.GetAllOrders).Methods("GET")
	s.Router.HandleFunc("/api/orders/{id}", s.OrderHandler.UpdateOrder).Methods("PUT")
	s.Router.HandleFunc("/api/orders/{id}", s.OrderHandler.GetOrderById).Methods("GET")

	// media
	s.Router.HandleFunc("/api/media", s.UploadImage).Methods("POST")
	s.Router.PathPrefix("/media/").Handler(http.StripPrefix("/media", http.FileServer(http.Dir("./media"))))

	// payment
	s.Router.HandleFunc("/api/payment/{orderId}", middleware.Cors(s.PaymentHandler.ClientSecret)).Methods("GET")
	s.Router.HandleFunc("/api/payment/{orderId}", middleware.Cors(middleware.Jwt(s.PaymentHandler.CreatePayment))).Methods("POST", "OPTIONS")

}
