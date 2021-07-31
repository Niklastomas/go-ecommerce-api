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

type Categories struct {
	logger *log.Logger
	db     *gorm.DB
}

type CategoryHandler interface {
	GetAllCategories(w http.ResponseWriter, r *http.Request)
	GetCategoryById(w http.ResponseWriter, r *http.Request)
	CreateCategory(w http.ResponseWriter, r *http.Request)
	DeleteCatecory(w http.ResponseWriter, r *http.Request)
	UpdateCategory(w http.ResponseWriter, r *http.Request)
}

func NewCategories(l *log.Logger, db *gorm.DB) *Categories {
	return &Categories{l, db}
}

func (c *Categories) GetAllCategories(w http.ResponseWriter, r *http.Request) {
	categories, err := models.GetAllCategories(c.db)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	responses.JSON(w, r, categories, http.StatusOK)
}

func (c *Categories) GetCategoryById(w http.ResponseWriter, r *http.Request) {
	var category *models.Category
	var err error

	id := mux.Vars(r)["id"]

	category, err = category.GetById(c.db, id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	responses.JSON(w, r, category, http.StatusOK)
}

func (c *Categories) CreateCategory(w http.ResponseWriter, r *http.Request) {
	var category *models.Category

	err := json.NewDecoder(r.Body).Decode(&category)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	category, err = category.Create(c.db)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	responses.JSON(w, r, category, http.StatusOK)

}

func (c *Categories) DeleteCatecory(w http.ResponseWriter, r *http.Request) {
	var category *models.Category
	id := mux.Vars(r)["id"]

	err := category.Delete(c.db, id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	responses.JSON(w, r, category, http.StatusNoContent)
}

func (c *Categories) UpdateCategory(w http.ResponseWriter, r *http.Request) {
	var category *models.Category
	var err error

	err = json.NewDecoder(r.Body).Decode(&category)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	id := mux.Vars(r)["id"]
	category, err = category.Update(c.db, id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	responses.JSON(w, r, category, http.StatusOK)
}
