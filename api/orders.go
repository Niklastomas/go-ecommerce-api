package api

import (
	"encoding/json"
	"net/http"

	"github.com/niklastomas/go-ecommerce-api/models"
	"github.com/niklastomas/go-ecommerce-api/responses"
)

func (s *Server) CreateOrder(w http.ResponseWriter, r *http.Request) {
	var order *models.Order
	var err error

	err = json.NewDecoder(r.Body).Decode(&order)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	order, err = order.Create(s.DB)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	responses.JSON(w, r, order, http.StatusOK)

}

func (s *Server) GetAllOrders(w http.ResponseWriter, r *http.Request) {
	orders, err := models.GetAllOrders(s.DB)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	responses.JSON(w, r, orders, http.StatusOK)
}
