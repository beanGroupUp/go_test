package main

import (
	"fmt"
	"go_test/week4/4-8/model"
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
	fmt.Println("连接成功")
}

func main() {
	//第一种
	users := []model.User{
		{Name: "欢喜"},
		{Name: "面向加薪学习"},
		{Name: "Go语言几件一本通"},
		{Name: "Go语言微服务核心架构师22讲"},
		{Name: "从0到Go语言微服务架构师"},
	}
	//第一种批量添加
	//db.Create(users)

	//第二种批量添加
	//db.CreateInBatches(users, 100)

	//第三种批量添加(无字段填写，则不添加)
	db.Model(&model.User{}).Create([]map[string]interface{}{
		{"Name": "欢喜"},
		{"Name": "面向加薪学习"},
		{"Name": "Go语言几件一本通"},
		{"Name": "Go语言微服务核心架构师22讲"},
		{"Name": "从0到Go语言微服务架构师"},
	})

	for _, user := range users {
		fmt.Println(user.ID)
	}
}
