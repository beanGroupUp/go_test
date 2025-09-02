package model

import (
	"database/sql"

	"gorm.io/gorm"
)

type Product struct {
	gorm.Model
	Code  sql.NullString
	Price uint
}

type Food struct {
	FoodId uint   `gorm:"primarykey"`
	Name   string `gorm:"column:food_name;type:varchar(64);index:idx_food_name,unique"`
}
