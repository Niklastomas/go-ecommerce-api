package models

import (
	"fmt"
	"log"
	"strings"

	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

const (
	Customer = iota + 1
	Employee
	Admin
)

type User struct {
	gorm.Model
	FirstName string `json:"fname" gorm:"size:255"`
	LastName  string `json:"lname" gorm:"size:255"`
	Email     string `json:"email" gorm:"unique"`
	Password  string `json:"password" gorm:"size:255"`
	Role      int    `json:"role"`
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
	fmt.Println(u.Password)
	u.Password = string(hash)
	fmt.Println(u.Password)

	u.Email = strings.ToLower(u.Email)
	return nil

}

func (u *User) CheckPassword(providedPassword string) error {
	err := bcrypt.CompareHashAndPassword([]byte(u.Password), []byte(providedPassword))

	if err != nil {
		log.Println(err)
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
