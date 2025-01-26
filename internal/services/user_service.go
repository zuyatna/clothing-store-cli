package services

import (
	"clothing-pair-project/internal/models"
	"clothing-pair-project/internal/repository"
)

type UserService struct {
	userRepository repository.UserRepository
}

func NewUserService(userRepository repository.UserRepository) *UserService {
	return &UserService{userRepository: userRepository}
}

func (service *UserService) GetAllUsers() ([]models.User, error) {
	return service.userRepository.FindAll()
}

func (service *UserService) GetUserByUsername(username string) (models.User, error) {
	return service.userRepository.FindByUsername(username)
}

func (service *UserService) AddUser(user models.User) error {
	return service.userRepository.Add(user)
}

func (service *UserService) UpdateUser(user models.User) error {
	return service.userRepository.Update(user)
}

func (service *UserService) DeleteUser(username string) error {
	return service.userRepository.Delete(username)
}
