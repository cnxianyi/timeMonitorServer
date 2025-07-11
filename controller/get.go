package controller

import (
	"github.com/gin-gonic/gin"
	"time"
	"timeMonitorServer/models"
	"timeMonitorServer/types"
)

func All(c *gin.Context) {

	userName := c.Query("userName")
	date := c.Query("date")

	if date == "" {
		date = time.Now().Format("2006-01-02")
	} else {
		_, err := time.Parse("2006-01-02", date)
		if err != nil {
			c.JSON(400, gin.H{
				"msg": err.Error(),
			})
			return
		}
	}

	if userName == "" {
		c.JSON(400, gin.H{
			"msg": "userName is empty",
		})
		return
	}

	userId, err := models.FindUserIdByUserName(userName)
	if err != nil {
		c.JSON(400, gin.H{
			"code": 400,
			"msg":  "err user",
		})
		return
	}

	res := models.FindAllByDay(userId, date)

	var lastTime time.Time

	var processResponses []types.ProcessResponse = []types.ProcessResponse{}
	for _, pm := range res {
		var titleResponses []types.TitleResponse
		for _, tm := range pm.Titles {

			if tm.UpdateTime.After(lastTime) {
				lastTime = tm.UpdateTime
			}

			titleResponses = append(titleResponses, types.TitleResponse{
				Title: tm.Title,
				Time:  tm.Time,
			})
		}

		processResponses = append(processResponses, types.ProcessResponse{
			Process: pm.Process,
			Hour:    pm.Hour,
			Titles:  titleResponses,
		})
	}

	c.JSON(200, gin.H{
		"code":     200,
		"msg":      "ok",
		"data":     processResponses,
		"lastTime": lastTime,
	})
	return
}
