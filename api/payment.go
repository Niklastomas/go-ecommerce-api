package api

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/niklastomas/go-ecommerce-api/models"
	"github.com/niklastomas/go-ecommerce-api/responses"
	"github.com/stripe/stripe-go"
	"github.com/stripe/stripe-go/paymentintent"
	"gorm.io/gorm"
)

type Payments struct {
	logger *log.Logger
	db     *gorm.DB
}

type PaymentHandler interface {
	ClientSecret(w http.ResponseWriter, r *http.Request)
	CreatePayment(w http.ResponseWriter, r *http.Request)
}

type CheckoutData struct {
	ClientSecret string `json:"client_secret"`
}

func NewPayments(l *log.Logger, db *gorm.DB) *Payments {
	return &Payments{logger: l, db: db}
}

func (p *Payments) ClientSecret(w http.ResponseWriter, r *http.Request) {
	var order *models.Order
	var err error

	orderId := mux.Vars(r)["orderId"]
	order, err = order.GetById(p.db, orderId)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// Save key in env
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

	data := &CheckoutData{ClientSecret: pi.ClientSecret}

	responses.JSON(w, r, data, http.StatusOK)

}

func (p *Payments) CreatePayment(w http.ResponseWriter, r *http.Request) {
	var payment *models.Payment
	var order *models.Order
	var err error

	ctx := r.Context()
	userId := ctx.Value("userId").(int)
	orderId := mux.Vars(r)["orderId"]
	order, err = order.GetById(p.db, orderId)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	err = json.NewDecoder(r.Body).Decode(&payment)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	payment.UserId = userId
	payment.OrderId = int(order.ID)
	payment.Amount = order.Total
	payment.Status = "paid"

	payment, err = payment.Create(p.db)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	order.PaymentId = int(payment.ID)
	order.Update(p.db, orderId)
	responses.JSON(w, r, payment, http.StatusOK)

}
