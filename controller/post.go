package controller

import (
	"github.com/gin-gonic/gin"
	"timeMonitorServer/models"
	"timeMonitorServer/types"
)

func Upload(c *gin.Context) {

	var form []types.UploadForm

	err := c.ShouldBindJSON(&form)
	if err != nil {
		c.JSON(400, gin.H{
			"code": 400,
			"err":  err.Error(),
		})
		return
	}

	models.InsertAllProcessAndTitle(form)

	c.JSON(200, gin.H{
		"code": 200,
		"msg":  "ok",
	})
}
