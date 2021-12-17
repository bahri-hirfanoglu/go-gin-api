package routes

import (
	"net/http"

	"gin-framework-api/controllers"
	"gin-framework-api/services"

	"github.com/gin-gonic/gin"
)

var (
	UserService    services.UserService       = services.New()
	UserController controllers.UserController = controllers.New(UserService)
)

func Routes(route *gin.Engine) {
	user := route.Group("/users")
	{
		user.GET("/", func(c *gin.Context) {
			c.JSON(http.StatusOK, UserController.FindAll())
		})

		user.POST("/", func(c *gin.Context) {
			err := UserController.Save(c)
			if err != nil {
				c.JSON(http.StatusBadRequest, gin.H{
					"error": err.Error(),
				})
			} else {
				c.JSON(http.StatusOK, gin.H{
					"message": "User save successfully!",
				})
			}
		})
	}
}
