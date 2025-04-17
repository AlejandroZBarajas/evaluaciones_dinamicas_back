package userDomain

import (
	userEntity "evaluaciones/src/users/domain/entity"
)

type UserInterface interface {
	CreateUser(user *userEntity.UserEntity) error

	GetById(id int32) (*userEntity.UserEntity, error)

	GetByEmail(email string) (*userEntity.UserEntity, error)

	GetByMatricula(matricula string) (*userEntity.UserEntity, error)

	GetAll() ([]*userEntity.UserEntity, error)

	GetByRole(role string) ([]*userEntity.UserEntity, error)

	/* 	UpdateUser(user *userEntity.UserEntity) error

	   	DeleteUser(id int32) error */

	ExistsByEmail(email string) (bool, error)
	ExistsByMatricula(matricula string) (bool, error)
}
