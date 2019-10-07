package main

import (
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	//所有user开头的请求都会进入到这个函数中
	r.GET("/user/*action", func(c *gin.Context) {
		c.String(200, "hello world")
	})
	r.Run("localhost:8080")
}
