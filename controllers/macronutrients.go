package controllers

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/hyjaz/fitness-goal-tracker-server/models"
)

func AddMacroNutrient(c *gin.Context) {
	var user models.User
	var macroNutrientWithIDAsString models.MacroNutrientsWithIDAsString

	if err := c.ShouldBindQuery(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// We are deserializing it to DailyIntakeTimeAsString because BindJSON does not convert a unix timestamp as string to time.
	if err := c.ShouldBindJSON(&macroNutrientWithIDAsString); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	log.Println(macroNutrientWithIDAsString)
	log.Println(user)
	err := models.AddMacroNutrient(macroNutrientWithIDAsString, &user)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, user)
}
