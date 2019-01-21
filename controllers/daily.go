package controllers

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/hyjaz/fitness-goal-tracker-server/models"
	parse "github.com/hyjaz/fitness-goal-tracker-server/utility"
)

// AddDaily add a dailyIntake and returns all documents
func AddDaily(c *gin.Context) {
	uuid := c.Query("uuid")

	// We are deserializing it to DailyIntakeTimeAsString because BindJSON does not convert a unix timestamp as string to time.
	var dailyIntakeAsString models.DailyIntakeTimeAsString
	if err := c.ShouldBindJSON(&dailyIntakeAsString); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err})
		return
	}
	log.Println(dailyIntakeAsString.ID)
	//Once we have the unit timestamp as string we parse it
	date := parse.ConvertUnixTimestampToTime(dailyIntakeAsString.Date)
	var user models.User
	//then we pass in all necessary values to create a daily intake
	err := models.AddDailyIntake(uuid, dailyIntakeAsString.ID, date, dailyIntakeAsString.MacroNutrients, &user)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err})
		return
	}
	c.JSON(http.StatusOK, user)
}
