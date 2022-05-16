package service

import (
	"AirFile/model"
	"AirFile/utils"
	"github.com/gin-gonic/gin"
)

func Config(c *gin.Context) *model.Response {
	key := c.DefaultPostForm("key", "")
	response := model.NewResponse()
	response.Message = "success"
	response.Result = utils.GetConfig(key)
	return response
}
