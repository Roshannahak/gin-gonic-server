package controllers

import (
	"context"
	"gin_rest_api/config"
	"gin_rest_api/models"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

var userCollection = config.GetCollection(config.DB, "users")

func GetUsers(c *gin.Context) {

	var users []models.User

	result, err := userCollection.Find(context.TODO(), bson.M{})
	if err != nil {
		log.Fatal("internal server error")
		return
	}

	for result.Next(context.TODO()) {
		var singleUser models.User
		err := result.Decode(&singleUser)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"success": false, "msg": "internal server error"})
			return
		}
		users = append(users, singleUser)
	}

	c.JSON(http.StatusOK, gin.H{"success": true, "data": users})
}

func CreateUser(c *gin.Context) {
	var user models.User

	if err := c.BindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"status": false, "msg": "bad request"})
		return
	}

	newUser := models.User{
		Id:        primitive.NewObjectID(),
		FirstName: user.FirstName,
		Age:       user.Age,
	}

	_, err := userCollection.InsertOne(context.TODO(), newUser)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"success": false, "msg": "internal server error"})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"success": true, "msg": "successfully created..", "data": newUser})
}

func DeleteUser(c *gin.Context) {
	userId := c.Param("userId")

	objId, _ := primitive.ObjectIDFromHex(userId)

	result, err := userCollection.DeleteOne(context.TODO(), bson.M{"_id": objId})

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"success": false})
		return
	}

	if result.DeletedCount < 1 {
		c.JSON(http.StatusNotFound, gin.H{"success": false, "msg": "user not found"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"success": true, "msg": "successfully deleted..", "data": result})
}

func SearchUser(c *gin.Context) {
	userId := c.Param("userId")

	objId, _ := primitive.ObjectIDFromHex(userId)

	var user models.User

	err := userCollection.FindOne(context.TODO(), bson.M{"_id": objId}).Decode(&user)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"success": false})
		return
	}

	c.JSON(http.StatusOK, gin.H{"success": true, "msg": "found", "data": user})
}

func UpdateUser(c *gin.Context) {
	userId := c.Param("userId")

	objId, _ := primitive.ObjectIDFromHex(userId)

	var user models.User

	if err := c.BindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"success": false, "msg": "bad request"})
		return
	}

	update := bson.M{"firstname": user.FirstName, "age": user.Age}

	result, err := userCollection.UpdateOne(context.TODO(), bson.M{"_id": objId}, bson.M{"$set": update})

	if result.MatchedCount < 1 {
		c.JSON(http.StatusNotFound, gin.H{"success": false, "msg": "user not found"})
		return
	}

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"success": false, "msg": "internal server error"})
		return
	}

	//get updated user object
	var userUpdated models.User
	if result.MatchedCount == 1 {
		err := userCollection.FindOne(context.TODO(), bson.M{"_id": objId}).Decode(&userUpdated)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"success": false, "msg": "internal server error"})
			return
		}
		c.JSON(http.StatusOK, gin.H{"success": true, "msg": "update successfully..", "data": userUpdated})
	}
}
