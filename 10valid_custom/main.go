package main

import (
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"gopkg.in/go-playground/validator.v8"
	"reflect"
	"time"
)

type Booking struct {
	CheckIn  time.Time `form:"check_in" binding:"required,bookabledate" time_format:"2006-01-02"`
	CheckOut time.Time `form:"check_out" binding:"required,gtfield=CheckIn" time_format:"2006-01-02"`
}

func bookAbleDate(v *validator.Validate, topStruct reflect.Value, currentStructOrField reflect.Value,
	field reflect.Value, filedType reflect.Type,
	fieldKind reflect.Kind, param string) bool {
	if data, ok := field.Interface().(time.Time); ok {
		today := time.Now()
		if data.Unix() > today.Unix() {
			return true
		}
	}
	return false
}

func main() {
	//自定义结构体验证
	r := gin.Default()

	if v, ok := binding.Validator.Engine().(*validator.Validate); ok {
		v.RegisterValidation("bookabledate", bookAbleDate)
	}

	r.GET("/bookable", func(c *gin.Context) {
		var b Booking
		if err := c.ShouldBind(&b); err != nil {
			c.String(500, "%v", err)
			return
		}
		c.String(200, "%v", b)
	})

	r.Run("localhost:8080")
}
