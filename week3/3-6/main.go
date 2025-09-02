package main

import (
	"fmt"
	"go_test/week3/3-6/model"
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	r.POST("/login", LoginHandler)
	r.Run(":9098")
}

func LoginHandler(context *gin.Context) {
	var login model.AccountLogin
	err := context.ShouldBind(&login)
	if err != nil {
		fmt.Println(err.Error())
		context.JSON(http.StatusOK, gin.H{
			"msg": "输入有误",
		})
		return
	}
	context.JSON(http.StatusOK, gin.H{
		"msg": fmt.Sprintf("%s-%s", login.AccountName, login.Password),
	})
}
