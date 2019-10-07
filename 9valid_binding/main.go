package main

import (
	"github.com/gin-gonic/gin"
)

type Person struct {
	Name string `form:"name" binding:"required,gt=10"`
	//Name     string    `form:"name" binding:"required|gt=10"`
	Address string `form:"address" binding:"required"`
}

func main() {
	//结构体验证
	r := gin.Default()

	r.GET("/testing", func(c *gin.Context) {
		var person Person
		if err := c.ShouldBind(&person); err != nil {
			c.String(500, "%v", err)
			return
		}
		c.String(200, "%v", person)
	})

	r.Run("localhost:8080")
}
