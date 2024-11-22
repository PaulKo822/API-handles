package handlers

import (
	"context"
	"log"
	"pet_project_1_etap/internal/models"
	"pet_project_1_etap/internal/userservice"
	"pet_project_1_etap/internal/web/users"
)

type UserHandler struct {
	Service *userservice.UserService
}

func NewUserHandler(service *userservice.UserService) *UserHandler {
	return &UserHandler{
		Service: service,
	}
}

// DeleteUsersId implements users.StrictServerInterface.
func (u *UserHandler) DeleteUsersID(_ context.Context, request users.DeleteUsersIdRequestObject) (users.DeleteUsersIdResponseObject, error) {
	userID := uint(request.Id)
	err := u.Service.DeleteUserByID(userID)
	if err != nil {
		log.Printf("Error deleting user with ID %d: %v", userID, err)
		return nil, err
	}

	response := users.DeleteUsersId200Response{
		Message: "The user was successfully deleted",
	}

	return response, nil
}

// PatchUsersId implements users.StrictServerInterface.
func (u *UserHandler) PatchUsersID(_ context.Context, request users.PatchUsersIdRequestObject) (users.PatchUsersIdResponseObject, error) {
	userID := uint(request.Id)

	userRequest := request.Body

	userToCreate := models.User{
		Email:    *userRequest.Email,
		Password: *userRequest.Password,
	}

	updatedUser, err := u.Service.PatchUserByID(userID, userToCreate)

	if err != nil {
		return nil, err
	}
	// создаем структуру респонс
	response := users.PatchUsersId200Response{
		ID:       &updatedUser.ID,
		Email:    &updatedUser.Email,
		Password: &updatedUser.Password,
	}
	// Просто возвращаем респонс!
	return response, nil
}

// GetUsers implements users.StrictServerInterface.
func (u *UserHandler) GetUsers(_ context.Context, _ users.GetUsersRequestObject) (users.GetUsersResponseObject, error) {
	allUsers, err := u.Service.GetUsers()
	if err != nil {
		return nil, err
	}

	response := users.GetUsers200JSONResponse{}

	for _, use := range allUsers {
		user := users.User{
			ID:       &use.ID,
			Email:    &use.Email,
			Password: &use.Password,
		}
		response = append(response, user)
	}

	return response, nil
}

// PostUsers implements users.StrictServerInterface.
func (u *UserHandler) PostUsers(_ context.Context, request users.PostUsersRequestObject) (users.PostUsersResponseObject, error) {
	userRequest := request.Body

	userCreate := models.User{
		Email:    *userRequest.Email,
		Password: *userRequest.Password,
	}
	createdUser, err := u.Service.PostUser(userCreate)

	if err != nil {
		return nil, err
	}

	response := users.PostUsers201JSONResponse{
		ID:       &createdUser.ID,
		Email:    &createdUser.Email,
		Password: &createdUser.Password,
	}

	return response, nil
}

// GetTasksByUserID implements users.StrictServerInterface.
func (u *UserHandler) GetTasksByUserID(_ context.Context, request users.GetUsersIdRequestObject) (users.GetUsersIdResponseObject, error) {
	userID := uint(request.Id)

	tasksUser, err := u.Service.GetTasksForUser(userID)

	if err != nil {
		return nil, err
	}

	response := users.GetUsersId200JSONResponse(tasksUser)

	return response, nil
}
