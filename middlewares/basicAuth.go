package middlewares

import "github.com/gin-gonic/gin"

func BasicAuth(user string, pass string) gin.HandlerFunc {
	return gin.BasicAuth(gin.Accounts{
		user: pass,
	})
}
