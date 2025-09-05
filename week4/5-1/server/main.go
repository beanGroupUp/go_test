package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	r.GET("/hello", func(c *gin.Context) {
		lesson := c.Query("lesson")
		c.JSON(http.StatusOK, gin.H{
			"msg": "我要学习：" + lesson,
		})
	})
	err := r.Run(":9098")
	if err != nil {
		return
	}
}
