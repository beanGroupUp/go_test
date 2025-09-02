package main

import (
	"fmt"
	"go_test/week4/4-3/model"
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

}

func main() {
	/*	db.Create(&model.Product{
		Code:  "D42",
		Price: 200,
	})*/

	var p model.Product
	/*	db.First(&p, 1)
		fmt.Println(p.Price, p.Code)*/

	/*	db.First(&p, "code=?", "D42")
		fmt.Println(p.Price)*/

	//db.Model(&p).Update("price", 300)
	//db.First("code=?", "D42")
	//fmt.Println(p.Price)
	/*	db.First(&p, 2) // 先查询出ID为2的记录，赋值给p
		//零值，不会更新
		db.Model(&p).Updates(model.Product{
			Code:  "FF4222",
			Price: 0,
		})
		//db.First(&p, 2) 是 GORM 中的一个查询方法
		db.First(&p, 2)
		fmt.Println(p.Code)
		fmt.Println(p.Price)*/

	//会更新零值
	/*	db.Model(&p).Where("id = ?", 1).Updates(map[string]interface{}{
		"Code":  "",
		"Price": 0,
	})*/
	//不会更新零值
	/*	db.Model(&p).Where("id = ?", 1).Updates(model.Product{
			Price: 0,
			Code:  "",
		})
		db.First(&p, 1)
		fmt.Println(p.Code)
		fmt.Println(p.Price)*/

	/**
	db.First(&p, 2) 是 GORM 中的一个查询方法，它的作用是：

	功能说明
	查询单条记录：从数据库中检索第一条匹配条件的记录

	按主键查询：当第二个参数是简单值时（如数字 2），GORM 会将其解释为主键值

	结果填充：将查询到的数据填充到传入的结构体变量中（这里是 &p）

	具体含义
	db：GORM 数据库连接实例

	First()：GORM 的查询方法，用于获取第一条记录

	&p：指向 model.Product 类型变量的指针，查询结果将存储在这里

	2：要查询的记录的主键值（ID = 2）
	*/

	//软删除
	db.Delete(&p, 10)
}
