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
		w.WriteHeader(http.StatusInternalServerError)
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
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	user := db.GetUserByID(ID)
	if user == nil {
		fmt.Printf("Error while fetching user with ID: %v", ID)
		w.WriteHeader(http.StatusNotFound)
		return
	}

	res, err := json.Marshal(user)
	if err != nil {
		fmt.Println("Error while marshalling user")
		w.WriteHeader(http.StatusInternalServerError)
		return
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
	if u == nil {
		fmt.Println("Error creating user")
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	res, err := json.Marshal(u)
	if err != nil {
		fmt.Println("Error while marshalling user")
		w.WriteHeader(http.StatusInternalServerError)
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
	if existingUser == nil {
		fmt.Printf("Error fetching user with ID: %v", updatedUser.ID)
		w.WriteHeader(http.StatusNotFound)
		return
	}

	u := db.UpdateUser(updatedUser, existingUser)
	if u == nil {
		w.WriteHeader(http.StatusBadRequest)
		w.WriteHeader(http.StatusInternalServerError)
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
	if u == nil {
		w.WriteHeader(http.StatusBadRequest)
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
