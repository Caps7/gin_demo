package main

import (
	"github.com/gin-gonic/autotls"
	"github.com/gin-gonic/gin"
)

func main() {
	//自动化证书配置
	r := gin.Default()

	r.GET("/test", func(c *gin.Context) {
		c.String(200, "test")
	})
	//r.Run("localhost:8080")
	autotls.Run(r, "www.itpp.tk")
}
