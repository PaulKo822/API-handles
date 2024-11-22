package handlers

import (
	"context"
	"pet_project_1_etap/internal/models"
	"pet_project_1_etap/internal/taskservice"
	"pet_project_1_etap/internal/web/tasks"
)

type TaskHandler struct {
	Service *taskservice.TaskService
}

// Нужна для создания структуры Handler на этапе инициализации приложения

func NewTaskHandler(service *taskservice.TaskService) *TaskHandler {
	return &TaskHandler{
		Service: service,
	}
}

// DeleteTasksId implements tasks.StrictServerInterface.
func (h *TaskHandler) DeleteTasksID(_ context.Context, request tasks.DeleteTasksIdRequestObject) (tasks.DeleteTasksIdResponseObject, error) {
	taskID := uint(request.Id)

	err := h.Service.DeleteTaskByID(taskID)

	if err != nil {
		return nil, err
	}

	response := tasks.DeleteTasksId200Response{
		Message: "The task was successfully deleted",
	}

	return response, nil
}

// PatchTasksId implements tasks.StrictServerInterface.
func (h *TaskHandler) PatchTasksID(_ context.Context, request tasks.PatchTasksIdRequestObject) (tasks.PatchTasksIdResponseObject, error) {
	taskID := uint(request.Id)

	// Распаковываем тело запроса напрямую, без декодера!
	taskRequest := request.Body
	// Обращаемся к сервису и создаем задачу
	taskToCreate := models.Task{
		Text:   *taskRequest.Task,
		UserID: *taskRequest.UserID,
		IsDone: *taskRequest.IsDone,
	}

	updatedTask, err := h.Service.UpdateTaskByID(taskID, taskToCreate)

	if err != nil {
		return nil, err
	}
	// создаем структуру респонс
	response := tasks.PatchTasksId200Response{
		ID:     &updatedTask.ID,
		Task:   &updatedTask.Text,
		UserID: &updatedTask.UserID,
		IsDone: &updatedTask.IsDone,
	}
	// Просто возвращаем респонс!
	return response, nil
}

// GetTasks implements tasks.StrictServerInterface.
func (h *TaskHandler) GetTasks(_ context.Context, _ tasks.GetTasksRequestObject) (tasks.GetTasksResponseObject, error) {
	// Получение всех задач из сервиса
	allTasks, err := h.Service.GetAllTasks()
	if err != nil {
		return nil, err
	}

	// Создаем переменную респон типа 200джейсонРеспонс
	// Которую мы потом передадим в качестве ответа
	response := tasks.GetTasks200JSONResponse{}

	// Заполняем слайс response всеми задачами из БД
	for _, tsk := range allTasks {
		task := tasks.Task{
			ID:     &tsk.ID,
			Task:   &tsk.Text,
			UserID: &tsk.UserID,
			IsDone: &tsk.IsDone,
		}
		response = append(response, task)
	}

	// САМОЕ ПРЕКРАСНОЕ. Возвращаем просто респонс и nil!
	return response, nil
}

// PostTasks implements tasks.StrictServerInterface.
func (h *TaskHandler) PostTasks(_ context.Context, request tasks.PostTasksRequestObject) (tasks.PostTasksResponseObject, error) {
	// Распаковываем тело запроса напрямую, без декодера!
	taskRequest := request.Body
	// Обращаемся к сервису и создаем задачу
	taskToCreate := models.Task{
		Text:   *taskRequest.Task,
		UserID: *taskRequest.UserID,
		IsDone: *taskRequest.IsDone,
	}
	createdTask, err := h.Service.CreateTask(taskToCreate)

	if err != nil {
		return nil, err
	}
	// создаем структуру респонс
	response := tasks.PostTasks201JSONResponse{
		ID:     &createdTask.ID,
		Task:   &createdTask.Text,
		UserID: &createdTask.UserID,
		IsDone: &createdTask.IsDone,
	}
	// Просто возвращаем респонс!
	return response, nil
}
