package controller

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/eli-bosch/remindAI/internal/db"
	"github.com/eli-bosch/remindAI/internal/models"
	"github.com/eli-bosch/remindAI/internal/utils"
	"github.com/gorilla/mux"
)

// GET methods
func GetAllUsers(w http.ResponseWriter, r *http.Request) {
	users := db.GetAllUsers()
	res, err := json.Marshal(users)
	if err != nil {
		fmt.Println("Error while marshaling users")
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func GetUserByID(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	userID := vars["user_id"]
	ID, err := strconv.ParseInt(userID, 0, 0)
	if err != nil {
		fmt.Println("Error while parsing ID")
		return
	}

	user := db.GetUserByID(ID)
	if user == nil {
		fmt.Printf("Error while getting user by ID %d: %v", ID, err)
	}

	res, err := json.Marshal(user)
	if err != nil {
		fmt.Println("Error while marshalling user")
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

// PUT methods
func CreateNewUser(w http.ResponseWriter, r *http.Request) {
	newUser := &models.User{}
	utils.ParseBody(r, newUser)

	u := db.CreateUser(newUser)

	res, err := json.Marshal(u)
	if err != nil {
		fmt.Println("Error while marshalling user")
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

// POST methods
func UpdateUser(w http.ResponseWriter, r *http.Request) {
	updatedUser := &models.User{}
	utils.ParseBody(r, updatedUser)

	existingUser := db.GetUserByID(int64(updatedUser.ID))
	u := db.UpdateUser(updatedUser, existingUser)
	if u == nil {
		return
	}

	res, err := json.Marshal(u)
	if err != nil {
		fmt.Println("Error while marshalling user")
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

// Delete methods
func DeleteUser(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	userID := vars["user_id"]
	ID, err := strconv.ParseInt(userID, 0, 0)
	if err != nil {
		fmt.Println("Error while parsing user ID")
		return
	}

	u := db.DeleteUser(ID)
	res, err := json.Marshal(u)
	if err != nil {
		fmt.Println("Error while marshalling user")
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}
