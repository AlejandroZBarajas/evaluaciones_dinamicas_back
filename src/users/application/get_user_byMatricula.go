package userApplication

import (
	userDomain "evaluaciones/src/users/domain"
	userEntity "evaluaciones/src/users/domain/entity"
)

type GetUserByMatricula struct {
	repo userDomain.UserInterface
}

func NewGetUserByMatricula(repo userDomain.UserInterface) *GetUserByMatricula {
	return &GetUserByMatricula{repo: repo}
}

func (gubm *GetUserByMatricula) Run(matricula string) (*userEntity.UserEntity, error) {
	return gubm.repo.GetByMatricula(matricula)
}
