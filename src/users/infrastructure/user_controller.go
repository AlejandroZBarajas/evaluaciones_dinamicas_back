package userInfrastructure

import (
	"encoding/json"
	userApplication "evaluaciones/src/users/application"
	userEntity "evaluaciones/src/users/domain/entity"
	"fmt"
	"log"
	"net/http"
	"strconv"
)

type UserController struct {
	CreateUseCase         *userApplication.CreateUser
	GetAllUseCase         *userApplication.GetAllUsers
	GetByEmailUseCase     *userApplication.GetUserByEmail
	GetByIdUseCase        *userApplication.GetUserById
	GetByMatriculaUseCase *userApplication.GetUserByMatricula
	GetByRoleUseCase      *userApplication.GetUsersByRole
}

func NewUserController(
	createUser *userApplication.CreateUser,
	getAllUsers *userApplication.GetAllUsers,
	getUserByEmail *userApplication.GetUserByEmail,
	getUserById *userApplication.GetUserById,
	getUserByMatricula *userApplication.GetUserByMatricula,
	getUsersByRole *userApplication.GetUsersByRole,
) *UserController {
	return &UserController{
		CreateUseCase:         createUser,
		GetAllUseCase:         getAllUsers,
		GetByEmailUseCase:     getUserByEmail,
		GetByIdUseCase:        getUserById,
		GetByMatriculaUseCase: getUserByMatricula,
		GetByRoleUseCase:      getUsersByRole,
	}
}

func (uc *UserController) HandleCreateUser(w http.ResponseWriter, r *http.Request) {
	var user userEntity.UserEntity
	if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}
	createdUser, err := uc.CreateUseCase.Run(user.Email, user.Matricula)
	if err != nil {
		log.Printf("❌ Error al crear usuario: %v\n", err)
		http.Error(w, "Could not create user (infra)", http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(createdUser)
}

func (uc *UserController) HandleGetAllUsers(w http.ResponseWriter, r *http.Request) {
	users, err := uc.GetAllUseCase.Run()
	if err != nil {
		http.Error(w, "No se pudo obtener usuarios", http.StatusInternalServerError)
		return
	}
	json.NewEncoder(w).Encode(users)
}

func (uc *UserController) HandleGetUserByEmail(w http.ResponseWriter, r *http.Request) {
	email := r.URL.Path[len("/users/email/"):]
	if email == "" {
		http.Error(w, "Missing email parameter", http.StatusBadRequest)
		return
	}
	user, err := uc.GetByEmailUseCase.Run(email)
	if err != nil {
		http.Error(w, "User not found", http.StatusNotFound)
		return
	}
	json.NewEncoder(w).Encode(user)
}

func (uc *UserController) HandleGetUserById(w http.ResponseWriter, r *http.Request) {
	idStr := r.URL.Query().Get("id")
	if idStr == "" {
		http.Error(w, "Missing id parameter", http.StatusBadRequest)
		return
	}
	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "Invalid id", http.StatusBadRequest)
		return
	}
	user, err := uc.GetByIdUseCase.Run(int32(id))
	if err != nil {
		http.Error(w, "User not found", http.StatusNotFound)
		return
	}
	json.NewEncoder(w).Encode(user)
}

func (uc *UserController) HandleGetUserByMatricula(w http.ResponseWriter, r *http.Request) {
	matricula := r.URL.Path[len("/users/matricula/"):]
	if matricula == "" {
		http.Error(w, "Missing matricula parameter", http.StatusBadRequest)
		return
	}
	user, err := uc.GetByMatriculaUseCase.Run(matricula)
	if err != nil {
		http.Error(w, "User not found", http.StatusNotFound)
		return
	}
	json.NewEncoder(w).Encode(user)
}

func (uc *UserController) HandleGetUsersByRole(w http.ResponseWriter, r *http.Request) {

	if r.Method != http.MethodPost {
		http.Error(w, "Método no permitido", http.StatusMethodNotAllowed)
		return
	}

	var body struct {
		RoleID int32 `json:"role_id"`
	}

	if err := json.NewDecoder(r.Body).Decode(&body); err != nil {
		fmt.Println("📥 Role ID recibido:", body.RoleID)

		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	users, err := uc.GetByRoleUseCase.Run(body.RoleID)
	if err != nil {
		http.Error(w, "Could not retrieve users by role", http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(users)

}
