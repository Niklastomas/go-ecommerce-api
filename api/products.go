package api

import (
	"encoding/json"
	"net/http"

	"github.com/niklastomas/go-ecommerce-api/models"
	"github.com/niklastomas/go-ecommerce-api/responses"
)

func (s *Server) CreateProduct(w http.ResponseWriter, r *http.Request) {
	var product *models.Product

	err := json.NewDecoder(r.Body).Decode(&product)
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

func (s *Server) GetAllProducts(w http.ResponseWriter, r *http.Request) {

	products, err := models.GetAllProducts(s.DB)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	responses.JSON(w, r, products, http.StatusOK)
}
