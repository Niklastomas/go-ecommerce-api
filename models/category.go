package models

import (
	"gorm.io/gorm"
)

type Category struct {
	gorm.Model
	Name string `json:"name"`
}

func (c *Category) Create(db *gorm.DB) (*Category, error) {
	err := db.Create(&c).Error
	if err != nil {
		return nil, err
	}
	return c, nil
}

func (c *Category) GetById(db *gorm.DB, id string) (*Category, error) {
	err := db.Find(&c, id).Error
	if err != nil {
		return nil, err
	}
	return c, nil
}

func GetAllCategories(db *gorm.DB) ([]Category, error) {
	var categories []Category
	result := db.Find(&categories)
	return categories, result.Error
}

func (c *Category) Delete(db *gorm.DB, id string) error {
	return db.Delete(&c).Error
}

func (c *Category) Update(db *gorm.DB, id string) (*Category, error) {
	var category *Category
	var err error

	err = db.Model(&c).Where("id = ?", id).Updates(&c).Error
	if err != nil {
		return nil, err
	}

	category, err = c.GetById(db, id)
	if err != nil {
		return nil, err
	}
	return category, nil

}
