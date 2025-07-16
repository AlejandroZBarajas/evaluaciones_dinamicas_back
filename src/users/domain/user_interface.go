package userDomain

import (
	userEntity "evaluaciones/src/users/domain/entity"
)

type UserInterface interface {
	Register(user *userEntity.UserEntity, plainPassword string) error

	GetCredentialsByEmail(email string) (*userEntity.UserEntity, error)

	CreateUser(user *userEntity.UserEntity) error

	GetById(id int32) (*userEntity.UserEntity, error)

	GetByEmail(email string) (*userEntity.UserEntity, error)

	GetByMatricula(matricula string) (*userEntity.UserEntity, error)

	GetAll() ([]*userEntity.UserEntity, error)

	GetByRole(role int32) ([]*userEntity.UserEntity, error)

	ExistsByEmail(email string) (bool, error)
	ExistsByMatricula(matricula string) (bool, error)
}
