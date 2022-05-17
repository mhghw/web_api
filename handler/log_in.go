package handler

import (
	"authorizor/store"
	"authorizor/utils"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

type LogInForm struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

func LoginHandler(c *gin.Context) {
	var form LogInForm
	if err := c.BindJSON(&form); err != nil {
		stdErr := fmt.Errorf("error binding json: %w", err)
		errorResponse(c, http.StatusBadRequest, stdErr)
		return
	}

	usr, err := store.UserFileStoreInstance.GetUser(form.Username)
	if err != nil {
		stdErr := fmt.Errorf("error finding user with Username: %v with error: %w", form.Username, err)
		errorResponse(c, http.StatusNotFound, stdErr)
		return
	}

	if usr.Password != utils.HashPassword(form.Password) {
		stdErr := fmt.Errorf("incorrect password")
		errorResponse(c, http.StatusForbidden, stdErr)
		return
	}

	token, err := utils.GenerateToken(usr.Username)
	if err != nil {
		errorResponse(c, 500, err)
		return
	}

	resp := map[string]interface{}{
		"token": token,
	}

	c.JSON(200, resp)

}
