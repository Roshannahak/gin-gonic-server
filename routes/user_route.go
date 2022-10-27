package routes

import (
	"gin_rest_api/controllers"

	"github.com/gin-gonic/gin"
)


func UserRoute(router *gin.Engine){
	router.GET("/users", controllers.GetUser)
	router.POST("/user", controllers.CreateUser)
	router.DELETE("/user/:userId", controllers.DeleteUser)
	router.GET("/user/:userId", controllers.SearchUser)
	router.PUT("/user/:userId", controllers.UpdateUser)
}