package models

import (
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	ID       uint   `gorm:"primary_key"`
	Email	 string `gorm:"unique"`
	Password string
	Role	 string `gorm:"default:'user'"`
}

func (user *User) HashPassword(password string) error {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	
	if err != nil {
		return err
	}
	user.Password = string(bytes)
	return nil
}

	func (user *User) CheckPassword(providedPassword string) error {
		return bcrypt.CompareHashAndPassword([]byte(user.Password),
	[]byte(providedPassword))
	}