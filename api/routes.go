package api

import (
	"fmt"
	"net/http"
)

func (s *Server) InitRoutes() {

	s.Router.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "Hello")
	}).Methods("GET")

	s.Router.HandleFunc("/api/users", s.UsersCREATE).Methods("POST")
	s.Router.HandleFunc("/api/users", s.UsersLIST).Methods("GET")
}
