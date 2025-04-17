package userEntity

type UserEntity struct {
	Id        int32  `json:"id"`
	Email     string `json:"email"`
	Matricula string `json:"matricula"`
	Role      string `json:"role"`
}

func CreateUser(email string, matricula string) *UserEntity {
	return &UserEntity{Email: email, Matricula: matricula}
}
