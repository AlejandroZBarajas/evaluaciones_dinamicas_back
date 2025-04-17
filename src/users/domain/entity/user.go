package userEntity

type UserEntity struct {
	Id        int32  `json:"id"`
	Email     string `json:"email"`
	Matricula string `json:"matricula"`
	RoleID    int32  `json:"role_id"`
}

func CreateUser(email string, matricula string) *UserEntity {
	return &UserEntity{Email: email, Matricula: matricula}
}
