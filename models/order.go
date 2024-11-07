package models

import "gorm.io/gorm"

type Order struct {
	gorm.Model
	ID          uint   `gorm:"primary_key"`
	UserID		uint	`gorm:"not null"`
	ProductID	uint	`gorm:"not null"`
	Quantity	int		`gorm:"not null"`
	Status		string	`gorm:"default:'Pending'"`
}