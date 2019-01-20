package controllers

import (
	"fmt"

	apiErrors "github.com/hyjaz/fitness-goal-tracker-server/errors"

	parse "github.com/hyjaz/fitness-goal-tracker-server/utility"

	"github.com/gin-gonic/gin"
	"github.com/hyjaz/fitness-goal-tracker-server/models"
)

// AddCycle adds a new cycle
func AddCycle(c *gin.Context) {
	m := make(map[string]string)
	m["uuid"] = c.Query("uuid")
	m["startTime"] = c.PostForm("startTime")
	m["endTime"] = c.PostForm("endTime")

	for key, value := range m {
		err := validate(key, value, c)
		if err != (apiErrors.APIErrors{}) {
			c.JSON(400, err)
			return
		}
	}

	startTime := parse.ConvertUnixTimestampeToTime(m["startTime"])
	endTime := parse.ConvertUnixTimestampeToTime(m["endTime"])
	models.AddCycle(m["uuid"], startTime, endTime)
}

func validate(key string, value string, c *gin.Context) apiErrors.APIErrors {
	if value == "" {
		err := fmt.Sprintf("Invalid value provided to %s", key)
		return apiErrors.APIErrors{Error: err}
	}
	return apiErrors.APIErrors{}
}
