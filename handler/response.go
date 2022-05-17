package handler

import "github.com/gin-gonic/gin"

func errorResponse(c *gin.Context, code int, err error) {
	c.JSON(
		code,
		map[string]interface{}{
			"error": err.Error(),
		},
	)
}
