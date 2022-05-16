package service

import (
	"AirFile/model"
	"AirFile/utils"
)

func Config(key string) *model.Response {
	response := model.NewResponse()
	response.Message = "success"
	response.Result = utils.GetConfig(key)
	return response
}