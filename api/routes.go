package api

import (
	"fmt"
	"net/http"
)

func (s *Server) InitRoutes() {

	s.Router.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "Hello")
	}).Methods("GET")

	// users
	s.Router.HandleFunc("/api/users", s.UsersLIST).Methods("GET")

	// auth
	s.Router.HandleFunc("/api/login", s.Login).Methods("POST")
	s.Router.HandleFunc("/api/register", s.Register).Methods("POST")

	//products
	s.Router.HandleFunc("/api/products", s.CreateProduct).Methods("POST")
	s.Router.HandleFunc("/api/products", s.GetAllProducts).Methods("GET")

}
