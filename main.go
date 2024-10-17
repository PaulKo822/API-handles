package main

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

var Message string

type requestBody struct {
	Message string `json:"message"`
}

func HelloHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello, %s", Message)
}

func MessageHandler(w http.ResponseWriter, r *http.Request) {
	var msg requestBody

	err := json.NewDecoder(r.Body).Decode(&msg)
	if err != nil {
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}
	defer r.Body.Close()

	Message = msg.Message

}

func main() {
	router := mux.NewRouter()
	router.HandleFunc("/api/hello", HelloHandler).Methods("GET")
	router.HandleFunc("/api/message", MessageHandler).Methods("POST")
	http.ListenAndServe(":8080", router)
}
