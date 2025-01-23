package cli

import (
	"clothing-pair-project/entity"
	"clothing-pair-project/service"
	"log"
)

func AddUser(userService service.UserService, username, email, password, role string) {
	user := entity.User{
		Username: username,
		Email:    email,
		Password: password,
		Role:     role,
	}

	err := userService.Add(user)
	if err != nil {
		log.Fatal("Failed to add user:", err.Error())
	}
	log.Println("Successfully added user:", user)
}

func FindAllUsers(userService service.UserService) {
	users, err := userService.FindAll()
	if err != nil {
		log.Fatal("Failed to find all users:", err.Error())
	}
	log.Println("Successfully found users:", users)
}

func FindUserByID(userService service.UserService, id int) {
	user, err := userService.FindByID(id)
	if err != nil {
		log.Fatal("Failed to find user by ID:", err.Error())
	}
	log.Println("Successfully found user by ID:", user)
}

func FindUserByUsername(userService service.UserService, username string) {
	user, err := userService.FindByUsername(username)
	if err != nil {
		log.Fatal("Failed to find user by username:", err.Error())
	}
	log.Println("Successfully found user by username:", user)
}

func UpdateUser(userService service.UserService, id int, username, email, password, role string) {
	user := entity.User{
		UserID:   id,
		Username: username,
		Email:    email,
		Password: password,
		Role:     role,
	}

	err := userService.Update(user)
	if err != nil {
		log.Fatal("Failed to update user:", err.Error())
	}
	log.Println("Successfully updated user:", user)
}

func DeleteUser(userService service.UserService, id int) {
	err := userService.Delete(id)
	if err != nil {
		log.Fatal("Failed to delete user:", err.Error())
	}
	log.Println("Successfully deleted user with ID:", id)
}
