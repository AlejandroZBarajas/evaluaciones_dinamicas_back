package userApplication

import (
	userDomain "evaluaciones/src/users/domain"
	userEntity "evaluaciones/src/users/domain/entity"
)

type GetAllUsers struct {
	repo userDomain.UserInterface
}

func NewGetAllUsers(repo userDomain.UserInterface) *GetAllUsers {
	return &GetAllUsers{repo: repo}
}

func (gau *GetAllUsers) Run() ([]*userEntity.UserEntity, error) {
	return gau.repo.GetAll()
}
