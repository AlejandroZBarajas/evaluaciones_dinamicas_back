package userApplication

import (
	userDomain "evaluaciones/src/users/domain"
	userEntity "evaluaciones/src/users/domain/entity"
	"fmt"
)

type CreateUser struct {
	repo userDomain.UserInterface
}

func NewCreateUser(repo userDomain.UserInterface) *CreateUser {
	return &CreateUser{repo: repo}
}

func (cu *CreateUser) Run(email string, matricula string) (*userEntity.UserEntity, error) {
	user := userEntity.CreateUser(email, matricula)
	err := cu.repo.CreateUser(user)
	if err != nil {
		return nil, fmt.Errorf("error al crear usuario (application): %w", err)
	}
	return user, nil
}
