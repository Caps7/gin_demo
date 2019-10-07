package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"time"
)

type Person struct {
	Name     string    `form:"name"`
	Address  string    `form:"address"`
	Birthday time.Time `form:"birthday" time_format:"2006-01-02"`
}

func main() {
	r := gin.Default()

	r.GET("/testing", testing)
	r.POST("/testing", testing)

	r.Run("localhost:8080")
}

func testing(c *gin.Context) {
	var person Person
	//ShouldBind会根据content-type进行不同binding操作
	err := c.ShouldBind(&person)
	if err != nil {
		fmt.Println(err)
	}
	c.String(200, "%v", person)
}
