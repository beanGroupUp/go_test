package model

import (
	"gorm.io/gorm"
)

type Company struct {
	gorm.Model
	Name string
}

type Employer struct {
	gorm.Model
	Name       string
	CompanyID  int
	Company    Company
	CreditCard []CreditCard
}

type CreditCard struct {
	gorm.Model
	Number     string
	EmployerId uint
}
