package api

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/niklastomas/go-ecommerce-api/models"
	"github.com/niklastomas/go-ecommerce-api/responses"
	"gorm.io/gorm"
)

type Users struct {
	logger *log.Logger
	db     *gorm.DB
}

type UserHandler interface {
	GetUsers(w http.ResponseWriter, r *http.Request)
	GetUserById(w http.ResponseWriter, r *http.Request)
	UpdateUser(w http.ResponseWriter, r *http.Request)
}

func NewUsers(l *log.Logger, db *gorm.DB) *Users {
	return &Users{logger: l, db: db}
}

func (u *Users) GetUsers(w http.ResponseWriter, r *http.Request) {
	var err error
	var users []models.User

	users, err = models.GetAllUser(u.db)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
	responses.JSON(w, r, users, http.StatusOK)

}

func (u *Users) GetUserById(w http.ResponseWriter, r *http.Request) {
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

	user, err = user.GetByID(u.db, strconv.Itoa(userId))
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	responses.JSON(w, r, user, http.StatusOK)

}

func (u *Users) UpdateUser(w http.ResponseWriter, r *http.Request) {
	var user *models.User
	var err error
	id := mux.Vars(r)["id"]

	err = json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	user, err = user.Update(u.db, id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	responses.JSON(w, r, user, http.StatusOK)

}
