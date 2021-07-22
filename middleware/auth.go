package middleware

import (
	"context"
	"net/http"
	"os"
	"strings"
	"time"

	"github.com/niklastomas/go-ecommerce-api/auth"
)

func Jwt(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		authHeader := r.Header.Get("Authorization")
		tokenString := strings.Split(authHeader, "Bearer ")[1]

		jwtWrapper := &auth.JwtWrapper{
			SecretKey:       os.Getenv("JWT_SECRET"),
			Issuer:          "e-commerce",
			ExpirationHours: int64(12 * time.Hour),
		}

		claims, err := jwtWrapper.ValidateToken(tokenString)
		if err != nil {
			http.Error(w, "Unauthorized", http.StatusUnauthorized)
			return
		}
		ctx := context.WithValue(r.Context(), "userId", claims.UserId)
		next(w, r.WithContext(ctx))

	}
}
