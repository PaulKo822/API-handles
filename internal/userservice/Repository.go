package userservice

import (
	"pet_project_1_etap/internal/models"

	"gorm.io/gorm"
)

type UserRepository interface {
	GetUsers() ([]models.User, error)

	PostUser(task models.User) (models.User, error)

	PatchUserByID(id uint, task models.User) (models.User, error)

	DeleteUserByID(id uint) error

	GetTasksForUser(userID uint) ([]models.Task, error)
}

type userRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) UserRepository {
	return &userRepository{db: db}
}

func (r *userRepository) GetUsers() ([]models.User, error) {
	var users []models.User
	err := r.db.Find(&users).Error
	return users, err
}

func (r *userRepository) PostUser(user models.User) (models.User, error) {
	result := r.db.Create(&user)
	if result.Error != nil {
		return models.User{}, result.Error
	}
	return user, nil
}

func (r *userRepository) PatchUserByID(id uint, user models.User) (models.User, error) {
	var existingUser models.User
	if err := r.db.First(&existingUser, id).Error; err != nil {
		return models.User{}, err
	}

	existingUser.Email = user.Email

	existingUser.Password = user.Password

	if err := r.db.Save(&existingUser).Error; err != nil {
		return models.User{}, err
	}

	return existingUser, nil
}

func (r *userRepository) DeleteUserByID(id uint) error {
	if err := r.db.Unscoped().Delete(&models.User{}, id).Error; err != nil {
		return err
	}
	return nil
}

func (r *userRepository) GetTasksForUser(userID uint) ([]models.Task, error) {
	var tasks []models.Task
	if err := r.db.Preload("User").Where("user_id = ? AND deleted_at IS NULL", userID).Find(&tasks).Error; err != nil {
		return nil, err
	}
	return tasks, nil
}
