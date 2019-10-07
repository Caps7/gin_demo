package main

import (
	"github.com/gin-gonic/gin"
)

func main() {
	//get获取参数
	//localhost:8080/test?first_name=Tom
	r := gin.Default()

	r.GET("/test", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"firstName": c.Query("first_name"),
			"lastName": c.DefaultQuery(
				"last_name",
				"lastDefaultName",
			),
		})
	})
	r.Run("localhost:8080")
}
