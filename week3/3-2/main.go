package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func hello(c *gin.Context) {
	panic("程序出错了")
	c.JSON(http.StatusOK, gin.H{
		"msg": "面相加薪学习-从0到Go语言微服务架构师",
	})
}

/*func main() {
	r := gin.Default()
	//gin.New()
	r.GET("/ping", hello)
	r.Run(":9098")
}*/

func main() {
	//r := gin.Default()
	r := gin.New()
	r.GET("/ping", hello)
	r.Run(":9098")
}
