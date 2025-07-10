package controller

import (
	"github.com/gin-gonic/gin"
	"timeMonitorServer/models"
)

func All(c *gin.Context) {
	res := models.FindAllByDay()

	c.JSON(200, gin.H{
		"code": 200,
		"msg":  "ok",
		"data": res,
	})
}
