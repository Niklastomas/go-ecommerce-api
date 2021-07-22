package models

import (
	"strconv"

	"gorm.io/gorm"
)

type Order struct {
	gorm.Model
	UserId     int         `json:"user_id"`
	User       User        `json:"-"`
	OrderItems []OrderItem `json:"order_items" gorm:"foreignKey:OrderId"`
	Total      float64     `json:"total"`
	PaymentId  int         `json:"payment_id"`
	Payment    Payment     `json:"payment"`
}

type OrderItem struct {
	gorm.Model
	ProductId int     `json:"product_id"`
	Product   Product `json:"-"`
	Quantity  int     `json:"quantity"`
	OrderId   int     `json:"order_id"`
	Order     Order   `json:"-"`
}

func (o *Order) Create(db *gorm.DB) (*Order, error) {
	var err error

	err = o.CalculateTotalCost(db)
	if err != nil {
		return nil, err
	}

	err = db.Create(&o).Error
	return o, err
}

func GetAllOrders(db *gorm.DB) ([]Order, error) {
	var orders []Order

	err := db.Preload("OrderItems.Product").Preload("User").Preload("Payment").Find(&orders).Error
	if err != nil {
		return nil, err
	}
	return orders, err
}

func (o *Order) GetById(db *gorm.DB, id string) (*Order, error) {
	err := db.Find(&o, id).Error
	if err != nil {
		return nil, err
	}
	return o, nil

}

func (o *Order) Update(db *gorm.DB, id string) (*Order, error) {
	var order *Order
	var err error

	err = db.Model(&o).Where("id = ?", id).Updates(&o).Error
	if err != nil {
		return nil, err
	}

	order, err = o.GetById(db, id)
	if err != nil {
		return nil, err
	}
	return order, nil

}

func (o *Order) CalculateTotalCost(db *gorm.DB) error {
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
