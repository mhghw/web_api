package handler

import (
	"authorizor/store"
	"fmt"

	"github.com/gin-gonic/gin"
)

func GetUserDataHandler(c *gin.Context) {
	username := c.GetString("username")

	datas, err := store.DataStoreInstance.GetUserData(username)
	if err != nil {
		stdErr := fmt.Errorf("error reading user data: %w", err)
		errorResponse(c, 500, stdErr)
		return
	}
	if len(datas) == 0 {
		c.JSON(404, nil)
	}

	resps := make([]map[string]interface{}, 0)
	for _, data := range datas {
		resp := map[string]interface{}{
			"name": data.Name,
			"data": data.Data,
		}
		resps = append(resps, resp)
	}

	c.JSON(200, resps)
}
