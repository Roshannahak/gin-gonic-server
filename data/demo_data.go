package data

import "gin_rest_api/models"

var Users = []models.User{
	{Id: 1, FirstName: "roshan", Age: 22},
	{Id: 2, FirstName: "raja", Age: 23},
	{Id: 3, FirstName: "aman", Age: 25},
	{Id: 4, FirstName: "pragati", Age: 15},
}

func GetUsersData() []models.User{
	return Users
}