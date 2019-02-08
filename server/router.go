package server

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/hyjaz/fitness-goal-tracker-server/controllers"
)

func initRoute() *gin.Engine {
	r := gin.Default()
	r.Use(corsMiddleware())
	r.GET("/status", func(c *gin.Context) { c.JSON(http.StatusOK, gin.H{"status": "ok"}) })
	r.GET("/", controllers.GetUser)
	r.POST("/cycle", controllers.AddCycle)
	r.POST("/daily", controllers.AddDaily)
	r.POST("/macronutrients", controllers.AddMacroNutrient)
	return r
}

func corsMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS, GET, PUT")

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}

		c.Next()
	}
}
