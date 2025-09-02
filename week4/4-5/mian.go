package main

import (
	"database/sql"
	"fmt"
	"go_test/week4/4-5/model"
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

	db.AutoMigrate(&model.Food{})
}

func main() {
	/*	db.Create(&model.Product{
		Code:  sql.NullString{String: "D666", Valid: true},
		Price: 8888,
	})*/

	var p model.Product
	db.First(&p, 8)
	fmt.Println(p.Code)
	//当vaild为true时，更新零值；为false时，不更新零值
	db.Model(&p).Updates(model.Product{
		Code:  sql.NullString{"", false},
		Price: 9999,
	})

}
