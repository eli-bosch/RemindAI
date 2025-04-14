package db

import (
	"fmt"

	"github.com/eli-bosch/remindAI/config"
	"github.com/eli-bosch/remindAI/internal/models"
)

func CreateUser(u *models.User) *models.User {
	db := config.GetDB()
	err := db.Create(u).Error
	if err != nil {
		fmt.Printf("Error creating user: %v\n", err)
		return nil
	}

	return u
}

func GetAllUsers() []models.User {
	var users []models.User

	db := config.GetDB()
	err := db.Find(&users).Error
	if err != nil {
		fmt.Printf("Error finding users: %v\n", err)
	}

	return users
}

func GetUserByID(ID int64) *models.User {
	var user models.User

	db := config.GetDB()
	err := db.First(&user, ID).Error
	if err != nil {
		fmt.Printf("Error finding user with ID %d: %v\n", ID, err)
	}

	return &user
}

func UpdateUser(updatedUser *models.User, existingUser *models.User) *models.User {
	db := config.GetDB()
	err := db.Model(&existingUser).Update(map[string]interface{}{
		"username": updatedUser.Username,
		"password": updatedUser.Password,
		"phone":    updatedUser.Phone,
		"first":    updatedUser.First,
		"last":     updatedUser.Last,
		"address":  updatedUser.Address,
		"city":     updatedUser.City,
		"country":  updatedUser.Country,
	}).Error
	if err != nil {
		fmt.Printf("Error updating user: %v", err)
		return nil
	}

	return existingUser
}

func DeleteUser(ID int64) *models.User {
	user := GetUserByID(ID)
	if user == nil {
		return nil
	}

	db := config.GetDB()
	err := db.Delete(&user).Error
	if err != nil {
		fmt.Printf("Error deleting user with ID %d: %v\n", ID, err)
	}

	return user
}
