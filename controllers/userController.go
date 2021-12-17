package controllers

import (
	"gin-framework-api/entity"
	"gin-framework-api/services"

	"github.com/gin-gonic/gin"
)

type UserController interface {
	FindAll() []entity.Users
	Save(ctx *gin.Context) error
}

type controller struct {
	service services.UserService
}

func New(service services.UserService) UserController {
	return &controller{
		service: service,
	}
}

func (c *controller) Save(ctx *gin.Context) error {
	var user entity.Users
	err := ctx.ShouldBindJSON(&user)
	if err != nil {
		return err
	}
	c.service.Save(user)
	return nil
}

func (c *controller) FindAll() []entity.Users {
	return c.service.FindAll()
}
