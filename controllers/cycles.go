package controllers

import (
	"net/http"

	parse "github.com/hyjaz/fitness-goal-tracker-server/utility"

	"github.com/gin-gonic/gin"
	"github.com/hyjaz/fitness-goal-tracker-server/models"
)

// AddCycle adds a new cycle
func AddCycle(c *gin.Context) {
	uuid := c.Query("uuid")
	var cycleWithTimeAsString models.CycleWithTimeAsString
	if err := c.ShouldBindJSON(&cycleWithTimeAsString); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err})
		return
	}

	startTime := parse.ConvertUnixTimestampToTime(cycleWithTimeAsString.StartTime)
	endTime := parse.ConvertUnixTimestampToTime(cycleWithTimeAsString.EndTime)

	var user models.User
	err := models.AddCycle(uuid, startTime, endTime, &user)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err})
		return
	}
	c.JSON(http.StatusOK, user)
}
