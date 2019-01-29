package controllers

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/hyjaz/fitness-goal-tracker-server/models"
)

// GetUser adds the user if it hasn't been and return the users' information
func GetUser(c *gin.Context) {
	var user models.User
	log.Println("Inside controller GetUser")
	if err := c.ShouldBindQuery(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	err := models.GetUser(&user)
	if err != nil {
		c.JSON(404, err.Error())
	} else {
		c.JSON(200, user)
	}
}
