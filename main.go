package main

import (
	"gin_rest_api/config"
	"gin_rest_api/routes"

	"github.com/gin-gonic/gin"
)

func main() {

	router := gin.Default()
	gin.SetMode(gin.DebugMode)

	//connect database
	config.ConnectDB()

	routes.UserRoute(router)

	router.Run("localhost:3131")
}
