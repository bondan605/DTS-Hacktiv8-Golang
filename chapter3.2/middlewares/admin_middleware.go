package middlewares

import (
	model "middleware/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

func AdminMiddleware(c *gin.Context) {
	role, roleIsExist := c.Get("role")
	if !roleIsExist {
		c.AbortWithStatusJSON(http.StatusBadRequest, model.ErrorResponse{
			Err: "Unauthorized",
		})
		return
	}

	if role.(bool) == false {
		c.AbortWithStatusJSON(http.StatusBadRequest, model.ErrorResponse{
			Err: "Unauthorized",
		})
		return
	}

	c.Next()
}
