package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

// Структура для JSON
type InputData struct {
	Text string `json:"text"`
}

// GET
func HelloHandler(w http.ResponseWriter, r *http.Request) {

	var message []Message
	if err := DB.Find(&message).Error; err != nil {
		http.Error(w, "Error retrieving records", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(message)
}

// POST
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

// Patch / Put

func PutHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	idStr := vars["id"] // Получаем ID из параметров маршрута

	// Преобразование ID из строки в целое число
	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "Invalid ID", http.StatusBadRequest)
		return
	}

	var updatedMessage Message
	if err := json.NewDecoder(r.Body).Decode(&updatedMessage); err != nil {
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}

	// Установка ID для обновления
	updatedMessage.ID = uint(id)

	// Обновление записи в базе данных
	if err := DB.Save(&updatedMessage).Error; err != nil {
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(updatedMessage)
}

// DELETE

// Структура для ответа
type Response struct {
	Message string `json:"message"`
}

func DeleteHandler(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)
	idStr := vars["id"] // Получаем ID из параметров маршрута

	// Преобразование ID из строки в целое число
	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "Invalid ID", http.StatusBadRequest)
		return
	}

	// Удаление записи по ID
	result := DB.Delete(&Message{}, id)
	if result.Error != nil {
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		return
	}

	if result.RowsAffected == 0 {
		w.WriteHeader(http.StatusNotFound)
		json.NewEncoder(w).Encode(Response{Message: "ID not found"})
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(Response{Message: "Record deleted successfully"})
}

func main() {

	// Вызываем метод InitDB() из файла db.go
	InitDB()

	// Автоматическая миграция модели Message
	DB.AutoMigrate(&Message{})

	router := mux.NewRouter()
	router.HandleFunc("/api/hello", HelloHandler).Methods("GET")
	router.HandleFunc("/api/message", MessageHandler).Methods("POST")
	router.HandleFunc("/update/{id}", PutHandler).Methods("Put")
	router.HandleFunc("/delete_id/{id}", DeleteHandler).Methods("DELETE")
	http.ListenAndServe(":8080", router)
}
