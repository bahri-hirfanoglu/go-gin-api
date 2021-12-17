package services

import "gin-framework-api/entity"

type UserService interface {
	Save(entity.Users) entity.Users
	FindAll() []entity.Users
}

type userService struct {
	users []entity.Users
}

func New() UserService {
	return &userService{}
}

func (service *userService) Save(user entity.Users) entity.Users {
	service.users = append(service.users, user)
	return user
}

func (service *userService) FindAll() []entity.Users {
	return service.users
}
