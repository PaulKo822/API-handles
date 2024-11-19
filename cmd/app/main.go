package main

import (
	"log"
	"pet_project_1_etap/internal/database"
	"pet_project_1_etap/internal/handlers"
	"pet_project_1_etap/internal/taskservice"
	"pet_project_1_etap/internal/userservice"
	"pet_project_1_etap/internal/web/tasks"
	"pet_project_1_etap/internal/web/users"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	database.InitDB()
	err := database.DB.AutoMigrate(&taskservice.Task{})
	if err != nil {
		log.Fatalf("Ошибка при миграции базы данных: %v", err)
	}

	//TASK
	tasksRepo := taskservice.NewTaskRepository(database.DB)
	tasksService := taskservice.NewService(tasksRepo)
	tasksHandler := handlers.NewTaskHandler(tasksService)

	//USER
	usersRepo := userservice.NewUserRepository(database.DB)
	usersService := userservice.NewService(usersRepo)
	usersHandler := handlers.NewUserHandler(usersService)

	// Инициализируем echo
	e := echo.New()

	// используем Logger и Recover
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	// Прикол для работы в echo. Передаем и регистрируем хендлер в echo
	strictTaskHandler := tasks.NewStrictHandler(tasksHandler, nil) // тут будет ошибка
	tasks.RegisterHandlers(e, strictTaskHandler)

	strictUserHandler := users.NewStrictHandler(usersHandler, nil) // тут будет ошибка
	users.RegisterHandlers(e, strictUserHandler)

	if err := e.Start(":8080"); err != nil {
		log.Fatalf("failed to start with err: %v", err)
	}
}
