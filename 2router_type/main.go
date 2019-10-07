package main

import (
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	r.GET("/get", func(c *gin.Context) {
		c.String(200, "GET")
	})

	r.POST("/post", func(c *gin.Context) {
		c.String(200, "POST")
	})

	r.Handle("DELETE", "/delete", func(c *gin.Context) {
		c.String(200, "DELETE")
	})

	r.Any("/any", func(c *gin.Context) {
		c.String(200, "any")
	})

	r.Run("localhost:8080")
}
