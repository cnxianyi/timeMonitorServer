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

	userId, err := models.FindUserIdByUserName(form[0].UserName)

	models.InsertAllProcessAndTitle(form, userId)

	limit := models.FindDailyLimit(form[0].UserName)

	all, err := models.ComputedAll(form[0].UserName)

	c.JSON(200, gin.H{
		"code": 200,
		"msg":  "ok",
		"data": types.UploadResponse{
			Lave:   int(limit - all),
			Notice: "",
		},
	})
}

func EditTime(c *gin.Context) {
	var form types.UploadTimeForm

	err := c.ShouldBindJSON(&form)
	if err != nil {
		c.JSON(400, gin.H{
			"code": 400,
			"err":  err.Error(),
		})
		return
	}

	err = models.EditUserTime(form.Username, form.Password, form.Time, *form.Cycle)
	if err != nil {
		c.JSON(400, gin.H{
			"code": 400,
			"err":  err.Error(),
		})
		return
	}

	c.JSON(200, gin.H{
		"code": 200,
		"msg":  "ok",
	})

}
