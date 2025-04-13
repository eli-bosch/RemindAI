package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {

	r := mux.NewRouter()
	http.Handle("/", r)

	fmt.Println("Server is listening...")
	log.Fatal(http.ListenAndServe("localhost:9010", r))
}
