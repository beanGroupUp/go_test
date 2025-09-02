package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

//require github.com/gin-gonic/gin v1.10.1
//go mod tidy
//go get -u依赖包
//go mod init名称
//go get 依赖@版本
//go build

func main() {
	//依赖管理
	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"msg": "ok",
		})
	})
	err := r.Run()
	if err != nil {
		return
	}

}
