package server

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/hyjaz/fitness-goal-tracker-server/controllers"
)

func initRoute() *gin.Engine {
	r := gin.Default()
	r.GET("/status", func(c *gin.Context) { c.JSON(http.StatusOK, gin.H{"status": "ok"}) })
	r.GET("/", controllers.GetUser)
	r.POST("/cycle", controllers.AddCycle)
	r.POST("/daily", controllers.AddDaily)
	r.POST("/macronutrients", controllers.AddMacroNutrient)
	return r
}
