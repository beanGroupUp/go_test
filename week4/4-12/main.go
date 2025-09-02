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

func AddProducts() {
	p1 := model.Product{
		Code: sql.NullString{
			String: "Go语言几件一本通",
			Valid:  true,
		},
		Price: 99,
	}

	p2 := model.Product{
		Code: sql.NullString{
			String: "Go语言微服务架构核心22讲",
			Valid:  true,
		},
		Price: 99,
	}

	var productList []model.Product

	productList = append(productList, p1)
	productList = append(productList, p2)

	db.CreateInBatches(p1, 2)
}

func main() {
	//delete
	//soft delete delete_at

	//var p model.Product
	//db.First(&p, 8)
	//db.Delete(&p)

	db.Where("price=?", 777).Delete(&model.Product{})
}
