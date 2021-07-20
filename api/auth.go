package api

import (
	"encoding/json"
	"net/http"
	"os"
	"time"

	"github.com/niklastomas/go-ecommerce-api/auth"
	"github.com/niklastomas/go-ecommerce-api/models"
	"github.com/niklastomas/go-ecommerce-api/responses"
)

type LoginPayload struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type LoginResponse struct {
	Token string `json:"token"`
}

func (s *Server) Login(w http.ResponseWriter, r *http.Request) {
	var payload LoginPayload
	var user models.User
	var err error

	jwtWrapper := &auth.JwtWrapper{
		SecretKey:       os.Getenv("JWT_SECRET"),
		Issuer:          "e-commerce",
		ExpirationHours: int64(12 * time.Hour),
	}

	err = json.NewDecoder(r.Body).Decode(&payload)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	result := s.DB.Where("email = ?", payload.Email).Find(&user)
	if result.Error != nil {
		http.Error(w, result.Error.Error(), http.StatusBadRequest)
		return
	}

	err = user.CheckPassword(payload.Password)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	token, err := jwtWrapper.GenerateToken(int(user.ID))
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	tokenResponse := LoginResponse{Token: token}
	responses.JSON(w, r, tokenResponse, http.StatusOK)

}

func (s *Server) Register(w http.ResponseWriter, r *http.Request) {
	var err error
	var user *models.User

	err = json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	user, err = user.Create(s.DB)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	responses.JSON(w, r, user, http.StatusCreated)
}

// func CreateToken(userId uint) (string, error) {
// 	var err error
// 	claims := jwt.MapClaims{}
// 	claims["userId"] = userId
// 	claims["exp"] = time.Now().Add(time.Minute * 10).Unix()

// 	t := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
// 	token, err := t.SignedString([]byte(os.Getenv("JWT_SECRET")))
// 	if err != nil {
// 		return "", err
// 	}
// 	return token, nil
// }
