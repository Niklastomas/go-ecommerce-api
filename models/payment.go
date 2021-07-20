package models

import "gorm.io/gorm"

type PaymentStatus string

type Payment struct {
	gorm.Model
	OrderId int           `json:"order_id"`
	Amount  float64       `json:"amount"`
	Status  PaymentStatus `json:"status"`
}

func (p *Payment) Create(db *gorm.DB) (*Payment, error) {
	err := db.Create(&p).Error
	if err != nil {
		return nil, err
	}
	return p, nil
}
