package handlers

import (
	"context"
	"pet_project_1_etap/internal/taskService"
	"pet_project_1_etap/internal/web/tasks"
)

type Handler struct {
	Service *taskService.TaskService
}

// Нужна для создания структуры Handler на этапе инициализации приложения

func NewHandler(service *taskService.TaskService) *Handler {
	return &Handler{
		Service: service,
	}
}

// GetTasks implements tasks.StrictServerInterface.
func (h *Handler) GetTasks(ctx context.Context, request tasks.GetTasksRequestObject) (tasks.GetTasksResponseObject, error) {
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
			Id:     &tsk.ID,
			Task:   &tsk.Text,
			IsDone: &tsk.IsDone,
		}
		response = append(response, task)
	}

	// САМОЕ ПРЕКРАСНОЕ. Возвращаем просто респонс и nil!
	return response, nil
}

// PostTasks implements tasks.StrictServerInterface.
func (h *Handler) PostTasks(ctx context.Context, request tasks.PostTasksRequestObject) (tasks.PostTasksResponseObject, error) {
	// Распаковываем тело запроса напрямую, без декодера!
	taskRequest := request.Body
	// Обращаемся к сервису и создаем задачу
	taskToCreate := taskService.Task{
		Text:   *taskRequest.Task,
		IsDone: *taskRequest.IsDone,
	}
	createdTask, err := h.Service.CreateTask(taskToCreate)

	if err != nil {
		return nil, err
	}
	// создаем структуру респонс
	response := tasks.PostTasks201JSONResponse{
		Id:     &createdTask.ID,
		Task:   &createdTask.Text,
		IsDone: &createdTask.IsDone,
	}
	// Просто возвращаем респонс!
	return response, nil
}

// DeleteTasksId implements tasks.StrictServerInterface.
func (h *Handler) DeleteTasksId(ctx context.Context, request tasks.DeleteTasksIdRequestObject) (tasks.DeleteTasksIdResponseObject, error) {
	taskId := uint(request.Id)

	err := h.Service.DeleteTaskByID(taskId)

	if err != nil {
		return nil, err
	}
	response := tasks.DeleteTasksId205Response{
		Message: "The task was successfully deleted",
	}

	return response, nil
}

// PatchTasksId implements tasks.StrictServerInterface.
func (h *Handler) PatchTasksId(ctx context.Context, request tasks.PatchTasksIdRequestObject) (tasks.PatchTasksIdResponseObject, error) {

	taskId := uint(request.Id)

	// Распаковываем тело запроса напрямую, без декодера!
	taskRequest := request.Body
	// Обращаемся к сервису и создаем задачу
	taskToCreate := taskService.Task{
		Text:   *taskRequest.Task,
		IsDone: *taskRequest.IsDone,
	}

	updatedTask, err := h.Service.UpdateTaskByID(taskId, taskToCreate)

	if err != nil {
		return nil, err
	}
	// создаем структуру респонс
	response := tasks.PatchTasksId204Response{
		Id:     &updatedTask.ID,
		Task:   &updatedTask.Text,
		IsDone: &updatedTask.IsDone,
	}
	// Просто возвращаем респонс!
	return response, nil
}
