package main

import (
	"go_test/week3/3-4/model"
	"math/rand"
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	//商品
	//
	productGroup := r.Group("/product")
	{
		productGroup.GET("/detail", detailHandler)
		productGroup.POST("/add", addHandler)
		productGroup.POST("/add/json", addJsonHandler)
	}

	r.Run(":9098")
}

func addJsonHandler(context *gin.Context) {
	var p model.Product
	err := context.ShouldBindJSON(&p)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{
			"msg": "解析参数错误",
		})
		return
	}
	context.JSON(http.StatusOK, gin.H{
		"id":   p.ID,
		"name": p.Name,
	})
}

func addHandler(context *gin.Context) {
	id := rand.Intn(10000)
	name := context.DefaultPostForm("name", "formDefaultName")
	context.JSON(http.StatusOK, gin.H{
		"id":   id,
		"name": name,
	})
}

func detailHandler(c *gin.Context) {
	id := c.DefaultQuery("id", "0")
	name := c.DefaultQuery("name", "defaultName")
	c.JSON(http.StatusOK, gin.H{
		"id":   id,
		"name": name,
	})
}
