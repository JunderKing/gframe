package controller

import (
	"gframe/helper/http"
	"gframe/service"

	"github.com/gin-gonic/gin"
)

type Demo struct{}

func (demo *Demo) CreateUser(c *gin.Context) {
	var params struct {
		Name string `form:"name" binding:"required"`
		Age  int    `form:"age" binding:"required"`
		Sex  int    `form:"sex"`
	}
	http.Validate(c, &params)
	demoService := &service.Demo{}
	userModel := demoService.CreateUser(params.Name, params.Age, params.Sex)

	http.RespOk(map[string]interface{}{
		"userId": userModel.Id,
	})
}

func (demo *Demo) GetUserInfo(c *gin.Context) {
	var params struct {
		UserId int `form:"userId" binding:"required"`
	}
	http.Validate(c, &params)
	demoService := &service.Demo{}
	userModel := demoService.FindUserById(params.UserId)

	http.RespOk(userModel)
}

func (demo *Demo) UpdateUserInfo(c *gin.Context) {
	var params struct {
		UserId int    `form:"userId" binding:"required"`
		Name   string `form:"name" binding:"required"`
	}
	http.Validate(c, &params)
	demoService := &service.Demo{}
	demoService.UpdateUserName(params.UserId, params.Name)

	http.RespOk(nil)
}

func (demo *Demo) DeleteUser(c *gin.Context) {
	var params struct {
		UserId int `form:"userId" binding:"required"`
	}
	http.Validate(c, &params)
	demoService := &service.Demo{}
	demoService.DeleteUser(params.UserId)

	http.RespOk(nil)
}
