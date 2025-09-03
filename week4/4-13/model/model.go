package model

import "gorm.io/gorm"

type Employer struct {
	//点开里面有删除时间+创建时间+更新时间+id
	gorm.Model
	Name      string
	CompanyID int
	Company   Company
}

type Company struct {
	gorm.Model
	Name string
}
