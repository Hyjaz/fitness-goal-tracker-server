package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/hyjaz/fitness-goal-tracker-server/models"
)

func GetUser(c *gin.Context) {
	user, err := models.GetUser(c.Query("uuid"))
	if err != nil {
		c.JSON(404, err)
	} else {
		c.JSON(200, user)
	}
}
