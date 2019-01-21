package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/hyjaz/fitness-goal-tracker-server/models"
	parse "github.com/hyjaz/fitness-goal-tracker-server/utility"
)

func AddDaily(c *gin.Context) {
	uuid := c.Query("uuid")
	// We are deserializing it to DailyIntakeTimeAsString because BindJSON does not convert a unix timestamp as string to time.
	var dailyIntakeAsString models.DailyIntakeTimeAsString
	if err := c.ShouldBindJSON(&dailyIntakeAsString); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err})
		return
	}

	//Once we have the unit timestamp as string we parse it
	date := parse.ConvertUnixTimestampeToTime(dailyIntakeAsString.Date)

	//then we pass in all necessary values to create a daily intake
	models.AddDailyIntake(uuid, date, dailyIntakeAsString.MacroNutrients)
}
