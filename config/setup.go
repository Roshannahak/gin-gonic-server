package config

import (
	"context"
	"fmt"
	"log"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

const driver = "mongodb+srv://roshannahak:112233Raja@cluster0.sq55i.mongodb.net/?retryWrites=true&w=majority"

func ConnectDB() *mongo.Client{
	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI(driver))

	if err != nil{
		log.Fatal("database connection error: ", err)
	}

	fmt.Println("database connected..")

	return client
}

//client instance
var DB *mongo.Client = ConnectDB()

//get database collection
func GetCollection(client *mongo.Client, collectionName string) *mongo.Collection{
	collection := client.Database("goDatabase").Collection(collectionName)

	return collection
}