package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/hyjaz/fitness-goal-tracker-server/models"
)

// AddMacroNutrient adds a new macronutrient in daily intake
func AddMacroNutrient(c *gin.Context) {
	var user models.User
	var macroNutrientsWithCycleAndDailyIntakeID models.MacroNutrientsWithCycleAndDailyIntakeID

	if err := c.ShouldBindQuery(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// We are deserializing it to DailyIntakeTimeAsString because BindJSON does not convert a unix timestamp as string to time.
	if err := c.ShouldBindJSON(&macroNutrientsWithCycleAndDailyIntakeID); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err := models.AddMacroNutrient(macroNutrientsWithCycleAndDailyIntakeID, &user)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, user)
}
