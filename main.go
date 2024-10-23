package main

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

// Структура для JSON
type InputData struct {
	Text string `json:"text"`
}

func HelloHandler(w http.ResponseWriter, r *http.Request) {

	var message []Message
	if err := DB.Find(&message).Error; err != nil {
		http.Error(w, "Error retrieving records", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(message)
}

func MessageHandler(w http.ResponseWriter, r *http.Request) {
	// Декодирование JSON
	var inputData InputData
	err := json.NewDecoder(r.Body).Decode(&inputData)
	if err != nil {
		http.Error(w, fmt.Sprintf("Error decoding JSON: %v", err), http.StatusBadRequest)
		return
	}
	defer r.Body.Close()

	// Создание новой записи в базе данных
	message := Message{Text: inputData.Text}
	if err := DB.Create(&message).Error; err != nil {
		http.Error(w, fmt.Sprintf("Error saving message to database: %v", err), http.StatusInternalServerError)
		return
	}

	// Успешный ответ
	w.WriteHeader(http.StatusCreated)
	fmt.Fprintf(w, "Message saved successfully with ID: %d", message.ID)
}

func main() {

	// Вызываем метод InitDB() из файла db.go
	InitDB()

	// Автоматическая миграция модели Message
	DB.AutoMigrate(&Message{})

	router := mux.NewRouter()
	router.HandleFunc("/api/hello", HelloHandler).Methods("GET")
	router.HandleFunc("/api/message", MessageHandler).Methods("POST")
	http.ListenAndServe(":8080", router)
}
