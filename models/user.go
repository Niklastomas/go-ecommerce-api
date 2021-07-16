package models

import (
	"strings"

	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Name     string `json:"name" gorm:"size:255"`
	Email    string `json:"email" gorm:"unique"`
	Password string `json:"-" gorm:"size:255"`
}

func (u *User) Create(db *gorm.DB) (*User, error) {
	result := db.Create(&u)
	return u, result.Error
}

func (u *User) BeforeCreate(tx *gorm.DB) error {
	hash, err := bcrypt.GenerateFromPassword([]byte(u.Password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}
	u.Password = string(hash)
	u.Email = strings.ToLower(u.Email)
	return nil

}

func (u *User) CheckPassword(providedPassword string) error {
	err := bcrypt.CompareHashAndPassword([]byte("hejhej"), []byte("hejhej"))

	if err != nil {
		return err
	}

	return nil
}

func GetAllUser(db *gorm.DB) ([]User, error) {
	var err error
	users := []User{}
	err = db.Model(&User{}).Find(&users).Error
	if err != nil {
		return nil, err
	}
	return users, err
}
