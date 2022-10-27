package controllers

import (
	"net/http"
	"strconv"

	"gin_rest_api/data"
	"gin_rest_api/models"

	"github.com/gin-gonic/gin"
)

func GetUser(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"success": "true", "status": "200", "result": data.Users})
}

func CreateUser(c *gin.Context) {
	var user models.User

	c.BindJSON(&user)

	newUser := models.User{Id: user.Id, FirstName: user.FirstName, Age: user.Age}

	data.Users = append(data.Users, newUser)

	c.JSON(http.StatusCreated, gin.H{"success": true, "msg": "successfully created", "result": newUser})
}

func DeleteUser(c *gin.Context) {
	userId := c.Param("userId")

	id, _ := strconv.Atoi(userId)

	for index, v := range data.Users {
		if v.Id == id {
			data.Users = append(data.Users[:index], data.Users[index+1:]...)
			c.JSON(http.StatusOK, gin.H{"success": true, "msg": "deleted successfully..", "result": v})
			return
		}
	}
	c.JSON(http.StatusNotFound, gin.H{"success": false, "msg": "user not found"})
}

func SearchUser(c *gin.Context) {
	userId := c.Param("userId")

	id, _ := strconv.Atoi(userId)

	for _, v := range data.Users {
		if v.Id == id {
			c.JSON(http.StatusOK, gin.H{"success": true, "item": 1, "result": v})
			return
		}
	}
	c.JSON(http.StatusNotFound, gin.H{"success": false, "msg": "not found"})
}

func UpdateUser(c *gin.Context) {
	userId := c.Param("userId")

	var user models.User

	c.BindJSON(&user)

	id, _ := strconv.Atoi(userId)

	for index, v := range data.Users {
		if v.Id == id {
			data.Users = append(data.Users[:index], data.Users[index+1:]...)
			data.Users = append(data.Users, user)
			c.JSON(http.StatusOK, gin.H{"success": true, "msg": "update successfully...", "result": data.Users})
			return
		}
	}
	c.JSON(http.StatusNotFound, gin.H{"success": false, "msg": "not found"})
}
