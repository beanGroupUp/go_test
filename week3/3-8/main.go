package main

import (
	"fmt"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

func CustomMiddleWare(c *gin.Context) {
	now := time.Now()
	params := c.Params
	name, ok := params.Get("name")
	if !ok {
		//中间件，终端后续的逻辑处理（拦截）
		c.Abort()
		//return只能return中间键的后续逻辑，业务逻辑还是照常执行
		//return
	}
	fmt.Println("name:", name)
	c.Next()
	expired := time.Now().Sub(now)
	fmt.Println("expired:", expired)
}

func main() {
	r := gin.Default()
	r.Use(CustomMiddleWare)
	r.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"msg": "OK",
		})
	})
	r.Run(":9098")
}
