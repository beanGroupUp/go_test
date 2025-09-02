package main

import (
	"fmt"
	"time"

	"github.com/gin-gonic/gin"
)

func Cook() {
	now := time.Now()
	fmt.Println("开始做饭")
	time.Sleep(3 * time.Second)
	fmt.Println("做饭完了")
	expired := time.Now().Sub(now)
	fmt.Println(expired)
}

func Buy() {
	fmt.Println("去买菜...")
	time.Sleep(3 * time.Second)
	fmt.Println("买菜回来...")
}

func Cook2() {
	fmt.Println("开始做饭")
	time.Sleep(3 * time.Second)
	fmt.Println("饭做完了")
}

func Eat() {
	fmt.Println("吃饭")
	time.Sleep(3 * time.Second)
	fmt.Println("饭做完了")
}

func Wash() {
	fmt.Println("开始洗完")
	time.Sleep(3 * time.Second)
	fmt.Println("洗碗结束")
}

func CustomMiddleWare(c *gin.Context) {
	now := time.Now()
	fmt.Println(c.Params)
	c.Next()
	expired := time.Now().Sub(now)
	fmt.Println(expired)
}

func CustomMiddleWare2() gin.HandlerFunc {
	return func(c *gin.Context) {
		now := time.Now()
		c.Next()
		expired := time.Now().Sub(now)
		fmt.Println("耗时：%v\n", expired)
	}
}

func main() {
	//Cook()
	// 创建一个默认的 Gin 路由引擎（包含 Logger 和 Recovery 中间件）
	r := gin.Default()
	// 注册自定义中间件 CustomMiddleWare
	r.Use(CustomMiddleWare)
	// 定义 GET 路由，路径为 "/:id"
	r.GET("/:id", func(c *gin.Context) {
		Buy()
		Cook2()
		Eat()
		Wash()
	})
	// 启动服务器并在 9098 端口监听
	r.Run(":9098")
}
