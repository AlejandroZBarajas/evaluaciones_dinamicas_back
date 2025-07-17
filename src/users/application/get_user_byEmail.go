package userApplication

import (
	userDomain "evaluaciones/src/users/domain"
	userEntity "evaluaciones/src/users/domain/entity"
)

type GetUserByEmail struct {
	repo userDomain.UserInterface
}

func NewGetUserByEmail(repo userDomain.UserInterface) *GetUserByEmail {
	return &GetUserByEmail{repo: repo}
}

func (gube *GetUserByEmail) Run(email string) (*userEntity.UserEntity, error) {
	return gube.repo.GetByEmail(email)
}
