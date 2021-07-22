package models

import (
	"fmt"

	"gorm.io/gorm"
)

type PaymentStatus string

type Payment struct {
	gorm.Model
	OrderId int           `json:"order_id"`
	Order   *Order        `json:"order"`
	UserId  int           `json:"user_id"`
	User    User          `json:"user"`
	Amount  float64       `json:"amount"`
	Status  PaymentStatus `json:"status"`
}

func (p *Payment) Create(db *gorm.DB) (*Payment, error) {
	err := db.Create(&p).Error
	if err != nil {
		return nil, err
	}
	fmt.Println(p)
	return p, nil
}
