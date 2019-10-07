package main

import (
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	r.LoadHTMLGlob("template/*")

	r.GET("/index", func(c *gin.Context) {
		c.HTML(200, "index.html", gin.H{
			"title": "this is title",
		})
	})

	r.Run("localhost:8080")
}
