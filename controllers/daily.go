package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/hyjaz/fitness-goal-tracker-server/models"
	parse "github.com/hyjaz/fitness-goal-tracker-server/utility"
)

// AddDaily add a dailyIntake and returns all documents
func AddDaily(c *gin.Context) {
	var user models.User
	var dailyIntakeAsString models.DailyIntakeTimeAsString

	if err := c.ShouldBindQuery(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// We are deserializing it to DailyIntakeTimeAsString because BindJSON does not convert a unix timestamp as string to time.
	if err := c.ShouldBindJSON(&dailyIntakeAsString); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	date := parse.ConvertUnixTimestampToTime(dailyIntakeAsString.Date)
	err := models.AddDailyIntake(dailyIntakeAsString.ID, date, dailyIntakeAsString.MacroNutrients, &user)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, user)
}
