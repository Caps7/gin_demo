package main

import (
	"github.com/gin-gonic/gin"
	"io"
	"os"
)

func main() {
	//Default 会默认添加engine.Use(Logger(), Recovery())这两个中间件
	//r := gin.Default()

	r := gin.New()

	f, _ := os.Create("gin.log")
	gin.DefaultWriter = io.MultiWriter(f)
	gin.DefaultErrorWriter = io.MultiWriter(f)

	r.Use(gin.Logger()) //Logger是请求时在终端打印请求URL和耗时等的中间件

	r.GET("/test", func(c *gin.Context) {
		name := c.DefaultQuery("name", "default_name")
		c.String(200, "%s", name)
	})

	r.Run("localhost:8080")
}
