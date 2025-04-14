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

// Get methods
func GetAllReminders(w http.ResponseWriter, r *http.Request) {
	reminders := db.GetAllReminder()

	res, err := json.Marshal(reminders)
	if err != nil {
		fmt.Println("Error while marshalling reminders")
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func GetReminderByID(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	reminderID := vars["reminder_id"]
	ID, err := strconv.ParseInt(reminderID, 0, 0)
	if err != nil {
		fmt.Println("Error while parsing reminder ID")
		return
	}

	reminder := db.GetReminderByID(ID)
	res, err := json.Marshal(reminder)
	if err != nil {
		fmt.Println("Error while marshalling reminder")
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func GetReminderByUserID(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	userID := vars["user_id"]
	ID, err := strconv.ParseInt(userID, 0, 0)
	if err != nil {
		fmt.Println("Error while parsing reminder ID")
		return
	}

	reminders := db.GetReminderByUserID(ID)
	res, err := json.Marshal(reminders)
	if err != nil {
		fmt.Println("Error while marshalling reminders")
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

// Put methods
func CreateNewReminder(w http.ResponseWriter, r *http.Request) {
	newReminder := &models.Reminder{}
	utils.ParseBody(r, newReminder)

	reminder := db.CreateReminder(newReminder)

	res, err := json.Marshal(reminder)
	if err != nil {
		fmt.Println("Error while marshalling reminder")
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func UpdateReminder(w http.ResponseWriter, r *http.Request) {
	updatedReminder := &models.Reminder{}
	utils.ParseBody(r, updatedReminder)

	reminder := db.UpdateReminder(updatedReminder)

	res, err := json.Marshal(reminder)
	if err != nil {
		fmt.Println("Error while marshalling reminder")
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func DeleteReminder(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	reminderID := vars["reminder_id"]
	ID, err := strconv.ParseInt(reminderID, 0, 0)
	if err != nil {
		fmt.Println("Error while parsing reminder ID")
		return
	}

	reminder := db.DeleteReminder(ID)
	res, err := json.Marshal(reminder)
	if err != nil {
		fmt.Println("Error while marshaling reminder")
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}
