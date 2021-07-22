package api

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/niklastomas/go-ecommerce-api/models"
	"github.com/niklastomas/go-ecommerce-api/responses"
)

func (s *Server) CreateProduct(w http.ResponseWriter, r *http.Request) {
	var product *models.Product
	var err error

	err = json.NewDecoder(r.Body).Decode(&product)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	product, err = product.Create(s.DB)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
	}
	responses.JSON(w, r, product, http.StatusOK)

}

func (s *Server) GetProductById(w http.ResponseWriter, r *http.Request) {
	var product *models.Product
	var err error

	id := mux.Vars(r)["id"]

	product, err = product.GetById(s.DB, id)
	if err != nil {
		http.Error(w, fmt.Sprintf("No product was found with id %s", id), http.StatusBadRequest)
		return
	}
	responses.JSON(w, r, product, http.StatusOK)

}

func (s *Server) GetAllProducts(w http.ResponseWriter, r *http.Request) {
	products, err := models.GetAllProducts(s.DB)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	responses.JSON(w, r, products, http.StatusOK)
}

func (s *Server) UpdateProduct(w http.ResponseWriter, r *http.Request) {
	var product *models.Product
	var err error
	id := mux.Vars(r)["id"]

	err = json.NewDecoder(r.Body).Decode(&product)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	product, err = product.Update(s.DB, id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	responses.JSON(w, r, product, http.StatusOK)
}
