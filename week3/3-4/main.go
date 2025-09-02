package main

import (
	"fmt"
	"go_test/week3/3-4/model"
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	//商品
	productGroup := r.Group("/product")
	{
		productGroup.GET("/productList", productList)
		//productGroup.GET("/priduct/1", productDetail1)
		//productGroup.GET("/product/2", productDetail2)
		productGroup.GET("/:id/:name", productDetail3)
		productGroup.GET("file/*all", fileHandler)
	}
	r.Run(":9098")
}

func productDetail3(context *gin.Context) {
	var p model.Product
	err := context.ShouldBindUri(&p)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{
			"msg": "输入有误",
		})
	} else {
		context.JSON(http.StatusOK, gin.H{
			"msg": fmt.Sprintf("%d-%s", p.ID, p.Name),
		})
	}
}

func fileHandler(context *gin.Context) {
	all := context.Param("all")
	context.JSON(http.StatusOK, gin.H{
		"msg": all,
	})
}

func productDetail(context *gin.Context) {
	id := context.Param("id")
	name := context.Param("name")
	if id == "" || name == "" {
		context.JSON(http.StatusBadRequest, gin.H{
			"msg": "参数请求错误",
		})
	} else {
		context.JSON(http.StatusOK, gin.H{
			"msg": id + " - " + name,
		})
	}

}

func productDetail2(context *gin.Context) {

}

func productDetail1(context *gin.Context) {

}

func productList(context *gin.Context) {

}
