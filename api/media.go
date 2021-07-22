package api

import (
	"fmt"
	"net/http"

	"github.com/niklastomas/go-ecommerce-api/responses"
	"github.com/niklastomas/go-ecommerce-api/utils"
)

func (s *Server) UploadImage(w http.ResponseWriter, r *http.Request) {
	fileName, err := utils.UploadFile(r)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	fmt.Println(fileName)
	responses.JSON(w, r, fileName, http.StatusOK)
}
