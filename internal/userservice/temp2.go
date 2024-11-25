package userservice

import "pet_project_1_etap/internal/models"

type UserService struct {
	repo UserRepository
}

func NewService(repo UserRepository) *UserService {
	return &UserService{repo: repo}
}

func (s *UserService) GetUsers() ([]models.User, error) {
	return s.repo.GetUsers()
}

func (s *UserService) PostUser(user models.User) (models.User, error) {
	return s.repo.PostUser(user)
}

func (s *UserService) PatchUserByID(id uint, user models.User) (models.User, error) {
	return s.repo.PatchUserByID(id, user)
}

func (s *UserService) DeleteUserByID(id uint) error {
	return s.repo.DeleteUserByID(id)
}

func (s *UserService) GetTasksForUser(userID uint) ([]models.Task, error) {
	return s.repo.GetTasksForUser(userID)
}
