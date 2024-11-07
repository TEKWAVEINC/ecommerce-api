package models

import "gorm.io/gorm"

type Product struct {
	gorm.Model
	ID       uint   `gorm:"primary_key"`
	Name		string
	Description string
	Price		float64
	Inventory	int
}