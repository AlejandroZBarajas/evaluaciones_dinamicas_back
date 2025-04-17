package userApplication

import (
	userDomain "evaluaciones/src/users/domain"
	userEntity "evaluaciones/src/users/domain/entity"
)

type GetUserById struct {
	repo userDomain.UserInterface
}

func NewGetUserById(repo userDomain.UserInterface) *GetUserById {
	return &GetUserById{repo: repo}
}

func (gubi *GetUserById) Run(id int32) (*userEntity.UserEntity, error) {
	return gubi.repo.GetById(id)
}
