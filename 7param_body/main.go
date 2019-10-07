package main

import (
	"bytes"
	"github.com/gin-gonic/gin"
	"io/ioutil"
)

func main() {
	//body获取参数
	r := gin.Default()

	r.POST("/test", func(c *gin.Context) {
		bodyBytes, _ := ioutil.ReadAll(c.Request.Body)
		//回存
		c.Request.Body = ioutil.NopCloser(bytes.NewBuffer(bodyBytes))
		c.String(200, string(bodyBytes))
	})
	r.Run("localhost:8080")
}
