package service

import (
	"AirFile/database"
	"AirFile/model"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"net/http"
)

func Auth(c *gin.Context) *model.Response {
	response := model.NewResponse()
	params := make(map[string]interface{})
	c.BindJSON(&params)
	uuidString := fmt.Sprintf("%v", params["uuid"])
	if uuidString == "" || uuidString == "<nil>" {
		uuidString = uuid.New().String()
	}
	var mUser []model.User
	database.DB.Where(&model.User{UUID: uuidString}).Find(&mUser)
	if len(mUser) <= 0 {
		common := model.Common{IpAddress: c.ClientIP()}
		mmUser := model.User{UUID: uuidString, Common: common}
		if err := database.DB.Create(&mmUser).Error; err != nil {
			fmt.Println(err)
			response.Code = 500
			response.Message = "插入新用户失败"
			c.JSON(http.StatusOK, response)
			return nil
		}
	}
	response.Message = "success"
	response.Result = uuidString
	return response
}
