package services

import "gin-framework-api/entity"

type UserService interface {
	Save(entity.Users) entity.Users
	FindAll() []entity.Users
}

type service struct {
	users []entity.Users
}

func New() UserService {
	return &service{}
}

func (service *service) Save(user entity.Users) entity.Users {
	service.users = append(service.users, user)
	return user
}

func (service *service) FindAll() []entity.Users {
	return service.users
}
