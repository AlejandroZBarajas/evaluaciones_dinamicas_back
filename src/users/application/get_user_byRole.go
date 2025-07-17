package userApplication

import (
	userDomain "evaluaciones/src/users/domain"
	userEntity "evaluaciones/src/users/domain/entity"
)

type GetUsersByRole struct {
	repo userDomain.UserInterface
}

func NewGetUsersByRole(repo userDomain.UserInterface) *GetUsersByRole {
	return &GetUsersByRole{repo: repo}
}

func (gubr *GetUsersByRole) Run(role int32) ([]*userEntity.UserEntity, error) {
	return gubr.repo.GetByRole(role)
}
