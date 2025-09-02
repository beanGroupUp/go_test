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
			ParameterizedQueries:      false,       // 设置为 false，在 SQL 日志中显示实际参数Don't include params in the SQL log
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
	//ErrMissingWhereClause
	//db.Model(&model.User{}).Where("name=?", "欢喜").Update("age", 18)
	//u1 := model.User{}
	//db.First(&u1, 7)
	//fmt.Println(u1.Name)
	//db.Model(&u1).Update("email", "abc@qq.com")

	//u2 := model.User{}
	//db.First(&u2, 9)
	//db.Model(&u2).Where("name=?", "Go语言几件一本通").Update("age", 18)

	//u3 := model.User{}
	//db.First(&u3, 9)
	//db.Model(&u3).Updates(model.User{Name: "吃货我来了", Age: 17})

	//u4 := model.User{}
	//db.First(&u4, 9)
	//db.Model(&u4).Updates(map[string]interface{}{
	//	"name": "面相加薪学习",
	//	"age":  18,
	//})

	//u5 := model.User{}
	//db.First(&u5, 9)
	//db.Model(&u5).Select("name").Updates(map[string]interface{}{
	//	"name": "吃货",
	//	"age":  17,
	//})

	//u6 := model.User{}
	//db.First(&u6, 9)
	////过滤name字段
	//db.Model(&u6).Omit("name").Updates(map[string]interface{}{
	//	"name": "吃货",
	//	"age":  17,
	//})

	u7 := model.User{}
	db.First(&u7, 9)
	//更新 select 里面的字段
	db.Model(&u7).Select("name", "age").Updates(map[string]interface{}{
		"name": "吃货",
		"age":  0,
	})
}
