package middleware

import (
	"authorizor/store"
	"authorizor/utils"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func Authenticator(c *gin.Context) {
	tokenStr := c.GetHeader("authorization")
	if tokenStr == "" {
		c.AbortWithStatus(http.StatusForbidden)
		return
	}

	username, err := utils.ParseToken(tokenStr)

	if err != nil {
		stdErr := fmt.Errorf("error parsing token: %w", err)
		c.AbortWithError(http.StatusBadRequest, stdErr)
		return
	}

	if _, err := store.UserFileStoreInstance.GetUser(username); err != nil {
		stdErr := fmt.Errorf("invalid token or user does not exist: %w", err)
		c.AbortWithError(http.StatusForbidden, stdErr)
		return
	}

	c.Set("username", username)
}
