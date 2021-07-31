package api

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/niklastomas/go-ecommerce-api/models"
	"github.com/niklastomas/go-ecommerce-api/responses"
	"gorm.io/gorm"
)

type Orders struct {
	logger *log.Logger
	db     *gorm.DB
}

type OrderHandler interface {
	CreateOrder(w http.ResponseWriter, r *http.Request)
	GetAllOrders(w http.ResponseWriter, r *http.Request)
	UpdateOrder(w http.ResponseWriter, r *http.Request)
	GetOrderById(w http.ResponseWriter, r *http.Request)
}

func NewOrders(l *log.Logger, db *gorm.DB) *Orders {
	return &Orders{logger: l, db: db}
}

func (o *Orders) CreateOrder(w http.ResponseWriter, r *http.Request) {
	var order *models.Order
	var err error

	err = json.NewDecoder(r.Body).Decode(&order)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	order, err = order.Create(o.db)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	responses.JSON(w, r, order, http.StatusOK)

}

func (o *Orders) GetAllOrders(w http.ResponseWriter, r *http.Request) {
	orders, err := models.GetAllOrders(o.db)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	responses.JSON(w, r, orders, http.StatusOK)
}

func (o *Orders) UpdateOrder(w http.ResponseWriter, r *http.Request) {
	var order *models.Order
	var err error
	id := mux.Vars(r)["id"]

	err = json.NewDecoder(r.Body).Decode(&order)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	order, err = order.Update(o.db, id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	responses.JSON(w, r, order, http.StatusOK)
}

func (o *Orders) GetOrderById(w http.ResponseWriter, r *http.Request) {
	var order *models.Order
	var err error
	id := mux.Vars(r)["id"]

	order, err = order.GetById(o.db, id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	responses.JSON(w, r, order, http.StatusOK)
}
