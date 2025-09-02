package model

import (
	"database/sql"
	"time"
)

type User struct {
	ID           uint
	Name         string
	Email        *string //string前面加*，此时也是可以吧引用变量的空值赋值进去的
	Age          int
	Birthday     *time.Time
	MemberNumber sql.NullString
	ActivatedAt  sql.NullTime
	CreatedAt    time.Time
	UpdatedAt    time.Time
}
