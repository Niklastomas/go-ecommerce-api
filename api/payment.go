package api

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/niklastomas/go-ecommerce-api/models"
	"github.com/niklastomas/go-ecommerce-api/responses"
	"github.com/stripe/stripe-go"
	"github.com/stripe/stripe-go/paymentintent"
)

// TODO
func (s *Server) Charge(w http.ResponseWriter, r *http.Request) {
	// var payment *models.Payment
	var order *models.Order
	var err error

	orderId := mux.Vars(r)["orderId"]
	order, err = order.GetById(s.DB, orderId)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	stripe.Key = "sk_test_51HW34HG0jNCpj1nF1L5pWrDE8wMzNouzh6vR1XKnkeCZgQepY3PP3F9axl1ca3Yt7g7wklvUsL6QRlruU1kRMocR0018fwawF6"

	params := &stripe.PaymentIntentParams{
		Amount:   stripe.Int64(int64(order.Total)),
		Currency: stripe.String(string(stripe.CurrencySEK)),
	}

	params.AddMetadata("integration_check", "accept_a_payment")

	pi, err := paymentintent.New(params)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	responses.JSON(w, r, pi, http.StatusOK)

}
