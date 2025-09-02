package main

import (
	"errors"
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
	//u1 := model.User{}
	//r := db.First(&u1, 2)
	//fmt.Println(&u1)
	//fmt.Println(u1.ID)
	//fmt.Println(r.Error)
	//fmt.Println(r.RowsAffected)
	//b := errors.Is(r.Error, gorm.ErrRecordNotFound)
	//if b {
	//	fmt.Println("查无此人")
	//} else {
	//	fmt.Println("他是" + u1.Name)
	//}

	//u2 := model.User{}
	//db.Take(&u2, 2)
	//fmt.Println(u2.ID)

	//u3 := model.User{}
	//r2 := db.First(&u3, 10)
	//fmt.Println(u3.Name)
	//fmt.Println(r2)

	//u3 := model.User{}
	//只会打印一行，并且优先打印id靠前的
	//r2 := db.First(&u3, 77, 9, 8)
	//fmt.Println(u3.ID)
	//fmt.Println(u3.Name)
	//fmt.Println(r2)

	//var users []model.User
	//r3 := db.Find(&users, []int{7, 8, 9})
	//fmt.Println(r3)
	//
	//for _, user := range users {
	//	fmt.Println(user.Name)
	//}

	//主键错写成字符串还能查询吗
	//u3 := model.User{}
	//r2 := db.Find(&u3, "7")
	//fmt.Println(u3.Name)
	//fmt.Println(r2)

	u3 := model.User{}
	r2 := db.Find(&u3, "golang")
	fmt.Println(u3.Name)
	fmt.Println(r2)
	b := errors.Is(r2.Error, gorm.ErrRecordNotFound)
	if b {
		fmt.Println("查无此人")
	} else {
		fmt.Println("他是" + u3.Name)
	}
}
