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

func (cu *CreateUser) Create(email string, matricula string) error {
	user := userEntity.CreateUser(email, matricula)

	err := cu.repo.CreateUser(user)

	if err != nil {
		return fmt.Errorf("error: %w", err)
	}
	return nil
}
