package controllers

import (
	"net/http"

	parse "github.com/hyjaz/fitness-goal-tracker-server/utility"

	"github.com/gin-gonic/gin"
	"github.com/hyjaz/fitness-goal-tracker-server/models"
)

// AddCycle adds a new cycle
func AddCycle(c *gin.Context) {
	// TODO: should maybe have a controllers data...
	// TODO: Currently using data model from the models....
	var user models.User
	var cycleWithTimeAsString models.CycleWithTimeAsString

	if err := c.ShouldBindQuery(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := c.ShouldBindJSON(&cycleWithTimeAsString); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	startTime := parse.ConvertUnixTimestampToTime(cycleWithTimeAsString.StartTime)
	endTime := parse.ConvertUnixTimestampToTime(cycleWithTimeAsString.EndTime)

	err := models.AddCycle(startTime, endTime, &user)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, user)
}
