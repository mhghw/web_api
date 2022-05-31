package main

import (
	"authorizor/handler"
	"authorizor/middleware"
	"authorizor/store"
	"authorizor/utils"
	"fmt"
	"log"
	"os"

	"github.com/gin-gonic/gin"
)

func main() {
	log.SetFlags(log.LstdFlags | log.Lshortfile)

	store.NewUserFileStore()
	store.NewDataStore()
	utils.InitSchemas()

	s := gin.Default()

	s.POST("/sign_up", handler.SignUpHandler)
	s.POST("/login", handler.LoginHandler)
	//s.POST("/send_validator", handler.SendValidatorHandler)

	authorized := s.Group("/", middleware.Authenticator)
	authorized.GET("/info", handler.GetInfoHandler)
	authorized.POST("/send_data/:name", handler.SendDataHandler)
	authorized.GET("/data/:name", handler.GetDataHandler)
	authorized.GET("/user_data", handler.GetUserDataHandler)

	var port string
	if port = os.Getenv("PORT"); port == "" {
		port = "8000"
	}
	s.Run(fmt.Sprintf(":%v", port))

	// authorized := s.Group("/", middleware.Authenticate)
	// authorized.POST("/login")
}
