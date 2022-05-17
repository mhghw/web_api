package handler

import (
	"authorizor/store"
	"fmt"

	"github.com/gin-gonic/gin"
)

func GetDataHandler(c *gin.Context) {
	username := c.GetString("username")
	dataName := c.Param("name")

	data, err := store.DataStoreInstance.GetData(username, dataName)
	if err != nil {
		stdErr := fmt.Errorf("error finding file:%v error: %w", dataName, err)
		errorResponse(c, 404, stdErr)
		return
	}

	resp := map[string]interface{}{
		"name": data.Name,
		"data": data.Data,
	}

	c.JSON(200, resp)
}
