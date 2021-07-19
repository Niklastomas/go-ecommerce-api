package models

import (
	"strconv"

	"gorm.io/gorm"
)

type PaymentStatus string

type Order struct {
	gorm.Model
	UserId        int           `json:"user_id"`
	User          User          `json:"user"`
	OrderItems    []OrderItem   `json:"order_items" gorm:"foreignKey:OrderId"`
	Total         float64       `json:"total"`
	PaymentStatus PaymentStatus `json:"payment_status"`
}

type OrderItem struct {
	gorm.Model
	ProductId int     `json:"product_id"`
	Product   Product `json:"product"`
	Quantity  int     `json:"quantity"`
	OrderId   int     `json:"order_id"`
	Order     Order   `json:"order"`
}

func (o *Order) Create(db *gorm.DB) (*Order, error) {
	var err error

	err = o.AutoCost(db)
	if err != nil {
		return nil, err
	}

	err = db.Create(&o).Error
	return o, err
}

func GetAllOrders(db *gorm.DB) ([]Order, error) {
	var orders []Order

	err := db.Preload("OrderItems.Product").Preload("User").Find(&orders).Error
	if err != nil {
		return nil, err
	}
	return orders, err
}

func (o *Order) AutoCost(db *gorm.DB) error {
	var product *Product
	var err error

	o.Total = float64(0)

	for _, item := range o.OrderItems {
		productId := strconv.Itoa(item.ProductId)
		product, err = product.GetById(db, productId)
		if err != nil {
			return err
		}

		o.Total += product.Price * float64(item.Quantity)

	}
	return nil
}
