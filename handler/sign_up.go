package handler

import (
	"authorizor/store"
	"authorizor/utils"
	"errors"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"gopkg.in/validator.v2"
)

type SignUpForm struct {
	Username        string `json:"user_name" validate:"min=3,max=40,regexp=^[a-zA-Z]*$"`
	FirstName       string `json:"first_name" validate:"nonzero"`
	LastName        string `json:"last_name" validate:"nonzero"`
	Password        string `json:"password" validate:"min=8"`
	ConfirmPassword string `json:"confirm_password" validate:"min=8"`
}

type SignUpResponse struct {
	Token string `json:"token"`
}

func SignUpHandler(c *gin.Context) {
	var form SignUpForm
	err := c.BindJSON(&form)
	if err != nil {
		stdErr := fmt.Errorf("error binding json in signup: %w\n", err)
		errorResponse(c, http.StatusBadRequest, stdErr)
		return
	}

	if err := validator.Validate(form); err != nil {
		stdErr := fmt.Errorf("error validating signup request: %w", err)
		errorResponse(c, http.StatusBadRequest, stdErr)
		return
	} else if form.Password != form.ConfirmPassword {
		errorResponse(c, http.StatusBadRequest, errors.New("passwords should match"))
		return
	}

	user := store.User{
		FirstName: form.FirstName,
		LastName:  form.LastName,
		Username:  form.Username,
		Password:  utils.HashPassword(form.Password),
	}

	err = store.UserFileStoreInstance.AddUser(user)
	if err != nil {
		stdErr := fmt.Errorf("error inserting user with username: %v with error: %v\n", user.Username, err)
		errorResponse(c, http.StatusInternalServerError, stdErr)
		return
	}

	token, err := utils.GenerateToken(form.Username)
	if err != nil {
		stdErr := fmt.Errorf("error creating token: %v\n", err)
		errorResponse(c, http.StatusInternalServerError, stdErr)
		return
	}

	// resp, err := json.Marshal(SignUpResponse{Token: token})
	// if err != nil {
	// 	log.Printf("error marshaling json: %v\n", err)
	// 	c.AbortWithStatus(http.StatusInternalServerError)
	// 	return
	// }
	resp := map[string]interface{}{
		"token": token,
	}

	c.JSON(200, resp)
}
