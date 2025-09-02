package main

import (
	"fmt"
	"go_test/week4/4-9/model"
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
	/*	u1 := model.User{}
		db.Where("name=?", "面向加薪学习").First(&u1)
		fmt.Println(u1.ID)*/

	/*	u2 := model.User{}
		db.Where(model.User{Name: "欢喜"}).First(&u2)
		fmt.Println(u2.ID)*/

	//var users []model.User
	////不等于
	//db.Where("name <> ?", "欢喜").Find(&users)
	//for _, user := range users {
	//	fmt.Println(user.ID)
	//}

	//var userList []model.User
	//db.Where(map[string]interface{}{"name": "欢喜"}).Find(&userList)
	//for _, user := range userList {
	//	fmt.Println(user)
	//}

	var userList2 []model.User
	db.Where(map[string]interface{}{
		"name": "欢喜",
		"age":  0,
	}).Find(&userList2)
	for _, user := range userList2 {
		fmt.Println(user)
	}
}
