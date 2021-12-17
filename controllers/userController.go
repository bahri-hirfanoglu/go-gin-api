package controllers

import (
	"gin-framework-api/entity"
	"gin-framework-api/services"
	"gin-framework-api/validators"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

type UserController interface {
	FindAll() []entity.Users
	Save(ctx *gin.Context) error
}

type controller struct {
	service services.UserService
}

var validate *validator.Validate

func New(service services.UserService) UserController {
	validate = validator.New()
	validate.RegisterValidation("is-image", validators.ValidatePictureUrl)
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
	err = validate.Struct(user)
	if err != nil {
		return err
	}
	c.service.Save(user)
	return nil
}

func (c *controller) FindAll() []entity.Users {
	return c.service.FindAll()
}
