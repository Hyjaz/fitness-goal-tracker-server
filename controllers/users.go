package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/hyjaz/fitness-goal-tracker-server/models"
)

// GetUser adds the user if it hasn't been and return the users' information
func GetUser(c *gin.Context) {
	uuid := c.Query("uuid")
	if uuid == "" {
		c.JSON(400, "Invalid uuid")
		return
	}
	var user models.User
	err := models.GetUser(uuid, &user)
	if err != nil {
		c.JSON(404, err)
	} else {
		c.JSON(200, user)
	}
}
