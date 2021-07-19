package api

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/niklastomas/go-ecommerce-api/models"
	"github.com/niklastomas/go-ecommerce-api/responses"
)

func (s *Server) GetAllCategories(w http.ResponseWriter, r *http.Request) {
	categories, err := models.GetAllCategories(s.DB)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	responses.JSON(w, r, categories, http.StatusOK)
}

func (s *Server) GetCategoryById(w http.ResponseWriter, r *http.Request) {
	var category *models.Category
	var err error

	id := mux.Vars(r)["id"]

	category, err = category.GetById(s.DB, id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	responses.JSON(w, r, category, http.StatusOK)
}

func (s *Server) CreateCategory(w http.ResponseWriter, r *http.Request) {
	var category *models.Category

	err := json.NewDecoder(r.Body).Decode(&category)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	category, err = category.Create(s.DB)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	responses.JSON(w, r, category, http.StatusOK)

}

func (s *Server) DeleteCatecory(w http.ResponseWriter, r *http.Request) {
	var category *models.Category
	id := mux.Vars(r)["id"]

	err := category.Delete(s.DB, id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	responses.JSON(w, r, category, http.StatusNoContent)
}

func (s *Server) UpdateCategory(w http.ResponseWriter, r *http.Request) {
	var category *models.Category
	var err error

	err = json.NewDecoder(r.Body).Decode(&category)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	id := mux.Vars(r)["id"]
	category, err = category.Update(s.DB, id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	responses.JSON(w, r, category, http.StatusOK)
}
