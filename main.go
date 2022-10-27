package main

import (
	"gin_rest_api/routes"
	"net/http"

	"github.com/gin-gonic/gin"
)

// controller
func homePage(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"data": "hello data"})
}

func main() {

	router := gin.Default()

	routes.UserRoute(router)

	router.Run("localhost:3131")
}
