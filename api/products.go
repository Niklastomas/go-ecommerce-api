package api

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/niklastomas/go-ecommerce-api/models"
	"github.com/niklastomas/go-ecommerce-api/responses"
	"gorm.io/gorm"
)

type Products struct {
	logger *log.Logger
	db     *gorm.DB
}

type ProductHandler interface {
	CreateProduct(w http.ResponseWriter, r *http.Request)
	GetProductById(w http.ResponseWriter, r *http.Request)
	GetAllProducts(w http.ResponseWriter, r *http.Request)
	UpdateProduct(w http.ResponseWriter, r *http.Request)
}

func NewProducts(l *log.Logger, db *gorm.DB) *Products {
	return &Products{l, db}
}

func (p *Products) CreateProduct(w http.ResponseWriter, r *http.Request) {
	var product *models.Product
	var err error

	err = json.NewDecoder(r.Body).Decode(&product)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	product, err = product.Create(p.db)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
	}
	responses.JSON(w, r, product, http.StatusOK)

}

func (p *Products) GetProductById(w http.ResponseWriter, r *http.Request) {
	var product *models.Product
	var err error

	id := mux.Vars(r)["id"]

	product, err = product.GetById(p.db, id)
	if err != nil {
		http.Error(w, fmt.Sprintf("No product was found with id %s", id), http.StatusBadRequest)
		return
	}
	responses.JSON(w, r, product, http.StatusOK)

}

func (p *Products) GetAllProducts(w http.ResponseWriter, r *http.Request) {
	products, err := models.GetAllProducts(p.db)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	responses.JSON(w, r, products, http.StatusOK)
}

func (p *Products) UpdateProduct(w http.ResponseWriter, r *http.Request) {
	var product *models.Product
	var err error
	id := mux.Vars(r)["id"]

	err = json.NewDecoder(r.Body).Decode(&product)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	product, err = product.Update(p.db, id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	responses.JSON(w, r, product, http.StatusOK)
}
