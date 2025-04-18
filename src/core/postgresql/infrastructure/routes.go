package database

import (
	questionInfrastructure "evaluaciones/src/question/infrastructure"
	userInfrastructure "evaluaciones/src/users/infrastructure"
	"net/http"
)

func RegisterUserRoutes(mux *http.ServeMux, controller *userInfrastructure.UserController) {
	mux.HandleFunc("/users", func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case http.MethodPost:
			controller.HandleCreateUser(w, r)
		case http.MethodGet:
			controller.HandleGetAllUsers(w, r)
		default:
			http.Error(w, "Método no permitido", http.StatusMethodNotAllowed)
		}
	})

	mux.HandleFunc("/users/id/", func(w http.ResponseWriter, r *http.Request) {
		if r.Method == http.MethodGet {
			controller.HandleGetUserById(w, r)
		} else {
			http.Error(w, "Método no permitido", http.StatusMethodNotAllowed)
		}
	})

	mux.HandleFunc("/users/email/", func(w http.ResponseWriter, r *http.Request) {
		if r.Method == http.MethodGet {
			controller.HandleGetUserByEmail(w, r)
		} else {
			http.Error(w, "Método no permitido", http.StatusMethodNotAllowed)
		}
	})

	mux.HandleFunc("/users/matricula/", func(w http.ResponseWriter, r *http.Request) {
		if r.Method == http.MethodGet {
			controller.HandleGetUserByMatricula(w, r)
		} else {
			http.Error(w, "Método no permitido", http.StatusMethodNotAllowed)
		}
	})

	mux.HandleFunc("/users/role", func(w http.ResponseWriter, r *http.Request) {
		if r.Method == http.MethodPost {
			controller.HandleGetUsersByRole(w, r)
		} else {
			http.Error(w, "Método no permitido", http.StatusMethodNotAllowed)
		}
	})

}
func RegisterQuestionRoutes(mux *http.ServeMux, controller *questionInfrastructure.QuestionController) {
	mux.HandleFunc("/questions", func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case http.MethodPost:
			controller.HandleCreateQuestion(w, r)
		case http.MethodPut:
			controller.HandleUpdateQuestion(w, r)
		case http.MethodDelete:
			controller.HandleDeleteQuestion(w, r)
		default:
			http.Error(w, "Método no permitido", http.StatusMethodNotAllowed)
		}
	})

	mux.HandleFunc("/questions/id", controller.HandleGetQuestionByID)
	mux.HandleFunc("/questions/exam", controller.HandleGetAllByExam)
	mux.HandleFunc("/questions/category", controller.HandleGetAllByCategory)
}
