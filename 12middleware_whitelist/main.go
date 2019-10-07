package main

import (
	"github.com/gin-gonic/gin"
)

func IPAuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		ipList := []string{
			"127.0.0.1",
		}
		flag := false
		clientIP := c.ClientIP()
		for _, host := range ipList {
			if clientIP == host {
				flag = true
				break
			}
		}
		if !flag {
			c.String(401, "%s", "clientip abanded")
		}
	}
}

func main() {
	r := gin.Default()

	r.Use(IPAuthMiddleware())

	r.GET("/test", func(c *gin.Context) {
		name := c.DefaultQuery("name", "default_name")
		c.String(200, "%s", name)
	})

	r.Run("localhost:8080")
}
