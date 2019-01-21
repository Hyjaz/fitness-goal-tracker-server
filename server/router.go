package server

import (
	"github.com/gin-gonic/gin"
	"github.com/hyjaz/fitness-goal-tracker-server/controllers"
)

func initRoute() *gin.Engine {
	r := gin.Default()
	r.GET("/", controllers.GetUser)
	r.POST("/cycle", controllers.AddCycle)
	r.POST("/daily", controllers.AddDaily)
	return r
}
