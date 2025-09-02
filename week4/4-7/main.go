package main

import (
	"database/sql"
	"fmt"
	"log"
	"os"
	"time"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var db *gorm.DB

func init() {
	var err error
	newLogger := logger.New(
		log.New(os.Stdout, "\r\n", log.LstdFlags), // io writer
		logger.Config{
			SlowThreshold:             time.Second, // Slow SQL threshold
			LogLevel:                  logger.Info, // Log level
			IgnoreRecordNotFoundError: true,        // Ignore ErrRecordNotFound error for logger
			ParameterizedQueries:      true,        // Don't include params in the SQL log
			Colorful:                  true,        // Disable color
		},
	)
	dsn := "root:123456@tcp(127.0.0.1:3306)/orm_test?charset=utf8mb4&parseTime=True&loc=Local"
	//注意赋值短变量声明 :=，这会在 init() 函数的作用域内创建一个新的局部变量 db，而不是赋值给包级别的全局变量 db。
	db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{
		Logger: newLogger,
	})
	if err != nil {
		panic(err)
	}
	fmt.Println("连接成功...")
	//db.AutoMigrate(&User{})
}

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

func main() {
	/*	now := time.Now()
		u1 := User{Name: "欢喜", Birthday: &now}
		result := db.Create(&u1)
		fmt.Println(u1.ID)
		fmt.Println(result.Error)
		fmt.Println(result.RowsAffected)*/

	//update可以更新零值（空值）
	//db.Model(&User{ID: 2}).Update("Name", "")
	//updates不可以更新零值（空值）
	//db.Model(&User{ID: 2}).Updates(User{Name: ""})
	//s := "abc@qq.com"
	//db.Model(&User{ID: 1}).Updates(User{Email: s})
	a := ""
	db.Model(&User{ID: 1}).Updates(User{Email: &a})
}
