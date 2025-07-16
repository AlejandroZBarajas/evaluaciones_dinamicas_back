package userApplication

import (
	"evaluaciones/src/core/auth"
	userDomain "evaluaciones/src/users/domain"
	"fmt"
)

type LoginUser struct {
	repo userDomain.UserInterface
}

func NewLoginUser(repo userDomain.UserInterface) *LoginUser { return &LoginUser{repo} }

type LoginOutput struct {
	AccessToken  string `json:"access_token"`
	RefreshToken string `json:"refresh_token"`
}

func (u *LoginUser) Run(email, password string) (*LoginOutput, error) {
	user, err := u.repo.GetCredentialsByEmail(email)
	if err != nil {
		return nil, fmt.Errorf("usuario no encontrado")
	}

	if !auth.CheckPasswordHash(password, user.PasswordHash) {
		return nil, fmt.Errorf("credenciales inválidas")
	}

	acc, ref, err := auth.GenerateTokens(user.Id, user.RoleID)
	if err != nil {
		return nil, err
	}

	return &LoginOutput{acc, ref}, nil
}
