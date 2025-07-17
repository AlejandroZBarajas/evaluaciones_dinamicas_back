package userApplication

import (
	userDomain "evaluaciones/src/users/domain"
	userEntity "evaluaciones/src/users/domain/entity"
)

type RegisterUser struct {
	repo userDomain.UserInterface
}

func NewRegisterUser(repo userDomain.UserInterface) *RegisterUser {
	return &RegisterUser{repo}
}

func (uc *RegisterUser) Run(user *userEntity.UserEntity, password string) error {
	return uc.repo.Register(user, password)
}
