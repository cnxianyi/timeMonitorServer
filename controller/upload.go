package controller

import (
	"fmt"
	"github.com/gin-gonic/gin"
	_type "timeMonitorServer/type"
)

func Upload(c *gin.Context) {

	form := _type.Times{}

	err := c.ShouldBindJSON(&form)
	if err != nil {
		c.JSON(400, gin.H{
			"code": 400,
			"err":  err.Error(),
		})
		return
	}

	for i := range form.Data {
		fmt.Println(form.Data[i])
	}

	c.JSON(200, gin.H{
		"code": 200,
		"msg":  "ok",
		"data": form.Data,
	})
}
