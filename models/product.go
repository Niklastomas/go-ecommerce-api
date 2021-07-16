package models

import "gorm.io/gorm"

type Product struct {
	gorm.Model
	Name         string  `json:"name"`
	Description  string  `json:"description"`
	Price        float64 `json:"price"`
	CountInStock int64   `json:"count_in_stock"`
	Image        string  `json:"image"`
	Rating       int64   `json:"rating"`
	UserId       *uint   `json:"user_id"`
	User         *User   `json:"user gorm:"foreignKey:UserID"`
}

func (p *Product) Create(db *gorm.DB) (*Product, error) {
	err := db.Create(&p).Error
	if err != nil {
		return nil, err
	}
	return p, nil

}

func (p *Product) Delete(db *gorm.DB, id uint) error {
	err := db.Delete(&p, id).Error
	if err != nil {
		return err
	}
	return nil
}

func (p *Product) Update(db *gorm.DB, product Product) error {
	err := db.Model(&p).Updates(product).Error
	if err != nil {
		return err
	}
	return nil
}

func GetAllProducts(db *gorm.DB) ([]*Product, error) {
	var products []*Product

	err := db.Find(&products).Error
	if err != nil {
		return nil, err
	}
	return products, err

}
