package database

import (
	categoryInfrastructure "evaluaciones/src/categories/infrastructure"
	questionInfrastructure "evaluaciones/src/question/infrastructure"
	studentExamInfrastructure "evaluaciones/src/student_exam/infrastrcuture"
	userInfrastructure "evaluaciones/src/users/infrastructure"

	examInfrastructure "evaluaciones/src/exam/infrastructure"
	resultsInfrastructure "evaluaciones/src/results/infrastructure"

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

	mux.HandleFunc("/auth/register", controller.HandleRegister)
	mux.HandleFunc("/auth/login", controller.HandleLogin)

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
func RegisterCategoryRoutes(mux *http.ServeMux, controller *categoryInfrastructure.CategoryController) {
	mux.HandleFunc("/categories", func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case http.MethodPost:
			controller.HandleCreateCategory(w, r)
		case http.MethodPut:
			controller.HandleUpdateCategory(w, r)
		case http.MethodDelete:
			controller.HandleDeleteCategory(w, r)
		default:
			http.Error(w, "Método no permitido", http.StatusMethodNotAllowed)
		}
	})

	mux.HandleFunc("/categories/id", func(w http.ResponseWriter, r *http.Request) {
		if r.Method == http.MethodPost {
			controller.HandleGetCategoryByID(w, r)
		} else {
			http.Error(w, "Método no permitido", http.StatusMethodNotAllowed)
		}
	})

	mux.HandleFunc("/categories/teacher", func(w http.ResponseWriter, r *http.Request) {
		if r.Method == http.MethodPost {
			controller.HandleGetCategoriesByTeacherID(w, r)
		} else {
			http.Error(w, "Método no permitido", http.StatusMethodNotAllowed)
		}
	})
}
func RegisterExamRoutes(mux *http.ServeMux, controller *examInfrastructure.ExamController) {
	mux.HandleFunc("/exams", func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case http.MethodPost:
			controller.HandleCreateExam(w, r)
		case http.MethodPut:
			controller.HandleUpdateExam(w, r)
		case http.MethodDelete:
			controller.HandleDeleteExam(w, r)
		default:
			http.Error(w, "Método no permitido", http.StatusMethodNotAllowed)
		}
	})

	mux.HandleFunc("/exams/id", func(w http.ResponseWriter, r *http.Request) {
		if r.Method == http.MethodPost {
			controller.HandleGetByID(w, r)
		} else {
			http.Error(w, "Método no permitido", http.StatusMethodNotAllowed)
		}
	})

	mux.HandleFunc("/exams/teacher", func(w http.ResponseWriter, r *http.Request) {
		if r.Method == http.MethodPost {
			controller.HandleGetAllByTeacher(w, r)
		} else {
			http.Error(w, "Método no permitido", http.StatusMethodNotAllowed)
		}
	})
}

func RegisterStudentExamRoutes(mux *http.ServeMux, controller *studentExamInfrastructure.StudentExamController) {

	mux.HandleFunc("/student-exams", func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case http.MethodPost:
			controller.HandleCreateStudentExam(w, r)
		default:
			http.Error(w, "Método no permitido", http.StatusMethodNotAllowed)
		}
	})

	mux.HandleFunc("/student-exams/", func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case http.MethodGet:
			controller.HandleGetStudentExamByID(w, r)
		case http.MethodDelete:
			controller.HandleDeleteStudentExam(w, r)
		default:
			http.Error(w, "Método no permitido", http.StatusMethodNotAllowed)
		}
	})

	mux.HandleFunc("/student-exams/by-exam", func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case http.MethodGet:
			controller.HandleGetAllByExamID(w, r)
		default:
			http.Error(w, "Método no permitido", http.StatusMethodNotAllowed)
		}
	})

	mux.HandleFunc("/evaluate_exam", func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case http.MethodPost:
			controller.HandleExamEvaluation(w, r)
		default:
			http.Error(w, "Método no permitido", http.StatusMethodNotAllowed)
		}
	})

	mux.HandleFunc("/student-exams/random", func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case http.MethodPost:
			controller.HandleGenerateRandomExam(w, r)
		default:
			http.Error(w, "Método no permitido", http.StatusMethodNotAllowed)
		}
	})

}

func RegisterResultRoutes(mux *http.ServeMux, controller *resultsInfrastructure.ResultsController) {
	mux.HandleFunc("/results", func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case http.MethodPost:
			controller.HandleCreateResult(w, r)
		case http.MethodDelete:
			controller.HandleDeleteAllResultsByExam(w, r)
		default:
			http.Error(w, "Método no permitido", http.StatusMethodNotAllowed)
		}
	})

	mux.HandleFunc("/results/by-exam", func(w http.ResponseWriter, r *http.Request) {
		if r.Method == http.MethodPost {
			controller.HandleGetAllResultsByExam(w, r)
		} else {
			http.Error(w, "Método no permitido", http.StatusMethodNotAllowed)
		}
	})
}
