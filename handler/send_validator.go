package handler

import (
	"authorizor/utils"
	"errors"
	"fmt"
	"io"
	"log"
	"os"

	"github.com/gin-gonic/gin"
)

var schemaDir = "../schema"

func SendValidatorHandler(c *gin.Context) {
	fileName := c.GetHeader("file_name")
	// body, err := ioutil.ReadAll(c.Request.Body)
	// if err != nil {
	// 	c.AbortWithError(500, fmt.Errorf("cannot read body: %w", err))
	// 	return
	// }

	err := os.MkdirAll(schemaDir, os.ModePerm)
	if err != nil && !os.IsExist(err) {
		errorResponse(c, 500, errors.New("cant make tmp dir : "+err.Error()))
		return
	}

	dst, err := os.Create(fmt.Sprintf("%v/%v", schemaDir, fileName))
	if err != nil {
		errorResponse(c, 500, errors.New("cant make dst files : "+err.Error()))
		return
	}
	defer dst.Close()

	_, err = io.Copy(dst, c.Request.Body)
	if err != nil {
		errorResponse(c, 500, errors.New("cant copy body : "+err.Error()))
		return
	}

	utils.LoadSchema(fileName)
	log.Println("schema loaded successfully")

	c.Status(200)

}
