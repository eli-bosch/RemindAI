package main

import (
	"fmt"
	"log"
	"net/http"

	_ "github.com/eli-bosch/remindAI/internal/db"
	"github.com/eli-bosch/remindAI/internal/routes"
	"github.com/gorilla/mux"
)

func main() {

	r := mux.NewRouter()
	http.Handle("/", r)

	routes.RegisterUserRoutes(r)
	routes.RegisterReminderRoutes(r)

	fmt.Println("Server is listening...")
	log.Fatal(http.ListenAndServe("localhost:9010", r))
}
