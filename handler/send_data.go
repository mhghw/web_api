package handler

import (
	"authorizor/store"
	"authorizor/utils"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/teris-io/shortid"
)

var tmpDir = "../tmp"
var dataStorageDir = "../store/storage/data"

func SendDataHandler(c *gin.Context) {
	username := c.GetString("username")

	user, err := store.UserFileStoreInstance.GetUser(username)
	if err != nil {
		errorResponse(c, 404, errors.New("user not found"))
		return
	}

	var body any

	err = c.BindJSON(&body)
	if err != nil {
		stdErr := fmt.Errorf("error binding json from body: %w", err)
		errorResponse(c, http.StatusBadRequest, stdErr)
		return
	}

	validator := c.GetHeader("file_type")

	b, err := json.Marshal(body)
	if err != nil {
		stdErr := fmt.Errorf("error json marshal: %w\n try using valid json\n", err)
		errorResponse(c, http.StatusInternalServerError, stdErr)
		return
	}

	if validator != "" {
		err = os.MkdirAll(tmpDir, os.ModePerm)
		if err != nil && !os.IsExist(err) {
			errorResponse(c, 500, errors.New("cant make tmp dir : "+err.Error()))
			return
		}

		tmpFilename := fmt.Sprintf("%v/%v.json", tmpDir, shortid.MustGenerate())
		err = ioutil.WriteFile(tmpFilename, b, 0644)
		if err != nil {
			stdErr := fmt.Errorf("error writing file to temp: %w", err)
			errorResponse(c, http.StatusInternalServerError, stdErr)
			return
		}

		valid, err := utils.ValidateData(validator, tmpFilename)
		if err != nil {
			os.Remove(tmpFilename)
			stdErr := fmt.Errorf("error validating file:%v with validator:%v error: %w", tmpFilename, validator, err)
			errorResponse(c, http.StatusInternalServerError, stdErr)
			return
		}

		if valid {
			dataname, err := store.DataStoreInstance.AddData(user.Username, body)
			if err != nil {
				os.Remove(tmpFilename)
				stdErr := fmt.Errorf("error storing data: %w", err)
				errorResponse(c, http.StatusInternalServerError, stdErr)
				return
			}
			os.Remove(tmpFilename)
			resp := map[string]interface{}{
				"file_name": dataname,
			}

			c.JSON(200, resp)
		} else {
			os.Remove(tmpFilename)
			stdErr := fmt.Errorf("json body is not valid with type: %v", validator)
			errorResponse(c, http.StatusBadRequest, stdErr)
			return
		}
	} else {
		stdErr := fmt.Errorf("need to provide file type")
		errorResponse(c, http.StatusBadRequest, stdErr)
		return
	}
}
