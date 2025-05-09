package db

import (
	"fmt"
	"time"

	"github.com/eli-bosch/remindAI/config"
	"github.com/eli-bosch/remindAI/internal/models"
)

func CreateReminder(r *models.Reminder) *models.Reminder { //FIX: Add API calls for google maps (travel time + remind at logic)
	db := config.GetDB()
	err := db.Create(r).Error
	if err != nil {
		fmt.Printf("Error creating reminder: %v\n", err)
	}

	return r
}

func GetAllReminder() []models.Reminder {
	var reminders []models.Reminder

	db := config.GetDB()
	err := db.Find(&reminders).Error
	if err != nil {
		fmt.Printf("Error finding reminders: %v\n", err)
	}

	return reminders
}

func GetReminderByID(ID int64) *models.Reminder {
	var reminder models.Reminder

	db := config.GetDB()
	err := db.First(&reminder, ID).Error
	if err != nil {
		fmt.Printf("Error finding reminder with ID %d: %v", ID, err)
	}

	return &reminder
}

func GetReminderByUserID(ID int64) *[]models.Reminder {
	var reminders []models.Reminder

	db := config.GetDB()
	err := db.Where("user_id=?", ID).Find(&reminders).Error
	if err != nil {
		fmt.Printf("Error finding reminder with user_id %d: %v", ID, err)
	}

	return &reminders
}

func GetReminderByRemindAt(currentTime time.Time) *[]models.Reminder {
	var reminder []models.Reminder

	db := config.GetDB()
	err := db.Where("remind_at <= ?", currentTime).Find(&reminder).Error
	if err != nil {
		fmt.Printf("Error finding current reminders")
	}

	return &reminder
}

func UpdateReminder(updatedReminder *models.Reminder) *models.Reminder {
	existingReminder := GetReminderByID(int64(updatedReminder.ID))
	if existingReminder == nil {
		return nil
	}

	db := config.GetDB()
	err := db.Model(&existingReminder).Update(map[string]interface{}{
		"title":       updatedReminder.Title,
		"description": updatedReminder.Description,
		"address":     updatedReminder.Address,
		"city":        updatedReminder.City,
		"zip":         updatedReminder.Zip,
		"country":     updatedReminder.City,
		"time":        updatedReminder.Time,
		"remind_at":   updatedReminder.RemindAt,
		"travel_time": updatedReminder.TravelTime,
		"repeating":   updatedReminder.Repeating,
		"end_date":    updatedReminder.EndDate,
	}).Error
	if err != nil {
		fmt.Printf("Error updating reminder: %v", err)
		return nil
	}

	return existingReminder
}

func DeleteReminder(ID int64) *models.Reminder {
	reminder := GetReminderByID(ID)
	if reminder == nil {
		return nil
	}

	db := config.GetDB()
	err := db.Delete(&reminder).Error
	if err != nil {
		fmt.Printf("Error deleting reminder with ID %d: %v\n", ID, err)
	}

	return reminder
}
