package api

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/niklastomas/go-ecommerce-api/models"
	"github.com/niklastomas/go-ecommerce-api/responses"
)

func (s *Server) UsersLIST(w http.ResponseWriter, r *http.Request) {
	var err error
	var users []models.User

	users, err = models.GetAllUser(s.DB)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
	responses.JSON(w, r, users, http.StatusOK)

}

func (s *Server) GetUserById(w http.ResponseWriter, r *http.Request) {
	var user *models.User
	var err error
	var userId int

	id := mux.Vars(r)["id"]
	if id == "me" {
		ctx := r.Context()
		userId = ctx.Value("userId").(int)
	} else {
		userId, err = strconv.Atoi(id)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
	}

	user, err = user.GetByID(s.DB, strconv.Itoa(userId))
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	responses.JSON(w, r, user, http.StatusOK)

}

func (s *Server) UpdateUser(w http.ResponseWriter, r *http.Request) {
	var user *models.User
	var err error
	id := mux.Vars(r)["id"]

	err = json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	user, err = user.Update(s.DB, id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	responses.JSON(w, r, user, http.StatusOK)

}
