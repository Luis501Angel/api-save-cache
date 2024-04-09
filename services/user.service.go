package services

import (
	"context"
	"encoding/json"
	"fill-asset-process-prototype/db"
	"fill-asset-process-prototype/models"
	"fmt"
)

func GetUsers() []models.User {
	var users []models.User
	err := db.Conn.Select(&users, "SELECT id, first_name, last_name, email FROM users.tb_users")
	if err != nil {
		fmt.Printf("Error fetching from database: %s\n", err.Error())
	}
	return users
}

func SaveUsersCache(users []models.User) bool {
	if len(users) > 0 {
		for _, element := range users {
			key := fmt.Sprintf("user:%s", element.Id)
			data, err := json.Marshal(element)

			if err != nil {
				fmt.Printf("Failed to marshall: %s", err.Error())
				return false
			}

			err = db.RedisClient.Set(context.Background(), key, data, 0).Err()
			if err != nil {
				fmt.Printf("Failed to set value in redis instance: %s", err.Error())
				return false
			}
		}
	}
	return true
}
