package routes

import (
	"github.com/eli-bosch/remindAI/internal/controller"
	"github.com/gorilla/mux"
)

var RegisterUserRoutes = func(router *mux.Router) {
	router.HandleFunc("/user/", controller.GetAllUsers).Methods("GET")
	router.HandleFunc("/user/{user_id}", controller.GetUserByID).Methods("GET")
	router.HandleFunc("/user/", controller.CreateNewUser).Methods("PUT")
	router.HandleFunc("/user/", controller.UpdateUser).Methods("POST")
	router.HandleFunc("/user/", controller.DeleteUser).Methods("DELETE")
}

var RegisterReminderRoutes = func(router *mux.Router) {
	router.HandleFunc("/reminder/", controller.GetAllReminders).Methods("GET")
	router.HandleFunc("/reminder/{reminder_id}", controller.GetReminderByID).Methods("GET")
	router.HandleFunc("/reminder/user/{user_id}", controller.GetReminderByUserID).Methods("GET")
	router.HandleFunc("/reminder/", controller.CreateNewReminder).Methods("PUT")
	router.HandleFunc("/reminder/", controller.UpdateReminder).Methods("POST")
	router.HandleFunc("/reminder/{reminder_id}", controller.DeleteReminder).Methods("DELETE")
}
