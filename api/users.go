package api

import (
	"net/http"

	"github.com/niklastomas/go-ecommerce-api/helpers"
	"github.com/niklastomas/go-ecommerce-api/models"
	"github.com/niklastomas/go-ecommerce-api/responses"
)

func (s *Server) UsersCREATE(w http.ResponseWriter, r *http.Request) {
	var err error
	var user *models.User

	err = helpers.ParseRequest(w, r, &user)
	if err != nil {
		return
	}
	user, err = user.Create(s.DB)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	responses.JSON(w, r, user)
}

func (s *Server) UsersLIST(w http.ResponseWriter, r *http.Request) {
	var err error
	var users []models.User
	users, err = models.GetAllUser(s.DB)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
	responses.JSON(w, r, users)

}
