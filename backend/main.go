package main

import (
	"backend/api/api/models"
	"backend/api/api/routers"
	"fmt"
)

func main() {
	models.Users = []models.User{
		{ID: 1, Username: "user1", Password: "password1", Email: "user1@gmail.com"},
		{ID: 2, Username: "user2", Password: "password2", Email: "user2@gmail.com"},
		{ID: 3, Username: "user3", Password: "password3", Email: "user3@gmail.com"},
	}

	fmt.Println("Iniciando servidor...")
	routers.HandleRequest()
}
