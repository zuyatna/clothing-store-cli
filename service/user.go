package service

import (
	"clothing-pair-project/entity"
	"clothing-pair-project/repository"
)

type UserService struct {
	userRepository repository.UserRepository
}

func NewUserService(userRepository repository.UserRepository) *UserService {
	return &UserService{userRepository}
}

func (service *UserService) FindAll() ([]entity.User, error) {
	users, err := service.userRepository.FindAll()
	if err != nil {
		return users, err
	}
	return users, nil
}

func (service *UserService) FindByID(userID int) (entity.User, error) {
	user, err := service.userRepository.FindByID(userID)
	if err != nil {
		return user, err
	}
	return user, nil
}

func (service *UserService) FindByUsername(username string) (entity.User, error) {
	user, err := service.userRepository.FindByUsername(username)
	if err != nil {
		return user, err
	}
	return user, nil
}

func (service *UserService) Add(user entity.User) error {
	err := service.userRepository.Add(user)
	if err != nil {
		return err
	}
	return nil
}

func (service *UserService) Update(user entity.User) error {
	err := service.userRepository.Update(user)
	if err != nil {
		return err
	}
	return nil
}

func (service *UserService) Delete(userID int) error {
	err := service.userRepository.Delete(userID)
	if err != nil {
		return err
	}
	return nil
}

func (service *UserService) ResetIncrement() error {
	err := service.userRepository.ResetIncrement()
	if err != nil {
		return err
	}
	return nil
}
