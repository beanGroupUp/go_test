package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func hello(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"msg": "面向加薪学习-从0到Go语言微服务架构师",
	})
}

func main() {
	r := gin.Default()
	//后面一段是匿名函数
	r.GET("/ping", hello)

	r.Run(":9098")
}
