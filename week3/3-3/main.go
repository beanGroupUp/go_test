package main

import (
	"go_test/week3/3-3/model"
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	//账户
	accountGroup := r.Group("/account")
	{
		accountGroup.GET("/list", accountList)
		accountGroup.POST("/add", addAccount)
	}

	//商品
	productGroup := r.Group("/product")
	{
		productGroup.GET("/productList", productList)
	}

	r.Run(":9098")
}

func productList(context *gin.Context) {

}

func addAccount(context *gin.Context) {

}

func accountList(context *gin.Context) {
	var accountList []model.Account
	a1 := model.Account{
		No:   1,
		Name: "老王",
	}
	accountList = append(accountList, a1)
	a2 := model.Account{
		No:   2,
		Name: "老郑",
	}
	accountList = append(accountList, a2)

	context.JSON(http.StatusOK, gin.H{
		"accountList": accountList,
	})
}
