package handler

import (
	"authorizor/store"
	"errors"

	"github.com/gin-gonic/gin"
)

type GetInfoResponse struct {
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Username  string `json:"username"`
}

func GetInfoHandler(c *gin.Context) {
	username := c.GetString("username")

	user, err := store.UserFileStoreInstance.GetUser(username)
	if err != nil {
		errorResponse(c, 404, errors.New("user not found"))
		return
	}

	resp := GetInfoResponse{
		Username:  user.Username,
		FirstName: user.FirstName,
		LastName:  user.LastName,
	}

	c.JSON(200, resp)

}
