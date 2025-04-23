package main

import (
	"log"
	"net/http"
	"os"

	database "evaluaciones/src/core/postgresql/infrastructure"

	userApplication "evaluaciones/src/users/application"
	userInfrastructure "evaluaciones/src/users/infrastructure"

	categoryApplication "evaluaciones/src/categories/application"
	categoryInfrastructure "evaluaciones/src/categories/infrastructure"

	examApplication "evaluaciones/src/exam/application"
	examInfrastructure "evaluaciones/src/exam/infrastructure"

	questionApplication "evaluaciones/src/question/application"
	questionInfrastructure "evaluaciones/src/question/infrastructure"

	studentExamApplication "evaluaciones/src/student_exam/application"
	studentExamInfrastructure "evaluaciones/src/student_exam/infrastrcuture"

	resultsApplication "evaluaciones/src/results/application"
	resultsInfrastructure "evaluaciones/src/results/infrastructure"

	"github.com/joho/godotenv"
)

func main() {
	if err := godotenv.Load(); err != nil {
		log.Fatal("❌ Error al cargar .env:", err)
	}

	if err := database.Connect(); err != nil {
		log.Fatal("❌ No se pudo conectar a la base de datos:", err)
	}
	db := database.GetDB()

	userRepo := userInfrastructure.NewUserRepository(db)

	createUser := userApplication.NewCreateUser(userRepo)
	getAllUsers := userApplication.NewGetAllUsers(userRepo)
	getUserByEmail := userApplication.NewGetUserByEmail(userRepo)
	getUserById := userApplication.NewGetUserById(userRepo)
	getUserByMatricula := userApplication.NewGetUserByMatricula(userRepo)
	getUsersByRole := userApplication.NewGetUsersByRole(userRepo)

	userController := userInfrastructure.NewUserController(
		createUser,
		getAllUsers,
		getUserByEmail,
		getUserById,
		getUserByMatricula,
		getUsersByRole,
	)

	categoryRepo := categoryInfrastructure.NewCategoryRepository(db)

	createCategory := categoryApplication.NewCreateCategory(categoryRepo)
	getCategoryById := categoryApplication.NewGetCategoryByID(categoryRepo)
	getCategoriesByTeacher := categoryApplication.NewGetAllCategoriesByTeacherID(categoryRepo)
	updateCategory := categoryApplication.NewUpdateCategory(categoryRepo)
	deleteCategory := categoryApplication.NewDeleteCategory(categoryRepo)

	categoryController := categoryInfrastructure.NewCategoryController(
		createCategory,
		getCategoryById,
		getCategoriesByTeacher,
		updateCategory,
		deleteCategory,
	)

	examRepo := examInfrastructure.NewExamRepository(db)

	createExam := examApplication.NewCreateExam(examRepo)
	getAllByTeacher := examApplication.NewGetAllExamsByTeacherID(examRepo)
	getExamByID := examApplication.NewGetExamByID(examRepo)
	updateExam := examApplication.NewUpdateExam(examRepo)
	deleteExam := examApplication.NewDeleteExam(examRepo)

	examController := examInfrastructure.NewExamController(
		createExam,
		getAllByTeacher,
		getExamByID,
		updateExam,
		deleteExam,
	)

	questionRepo := questionInfrastructure.NewQuestionRepository(db)

	createQuestion := questionApplication.NewCreateQuestion(questionRepo)
	getQuestionByID := questionApplication.NewGetQuestionByID(questionRepo)
	getAllByExam := questionApplication.NewGetAllQuestionsByExam(questionRepo)
	getAllByCategory := questionApplication.NewGetAllQuestionsByCategory(questionRepo)
	updateQuestion := questionApplication.NewUpdateQuestion(questionRepo)
	deleteQuestion := questionApplication.NewDeleteQuestion(questionRepo)

	questionController := questionInfrastructure.NewQuestionController(
		createQuestion,
		getQuestionByID,
		updateQuestion,
		deleteQuestion,
		getAllByExam,
		getAllByCategory,
	)

	studentExamRepo := studentExamInfrastructure.NewStudentExamRepository(db)

	createStudentExam := studentExamApplication.NewCreateStudentExam(studentExamRepo)
	getStudentExamByID := studentExamApplication.NewGetStudentExamByID(studentExamRepo)
	getAllStudentExams := studentExamApplication.NewGetAllByExamID(studentExamRepo)
	deleteStudentExam := studentExamApplication.NewDeleteStudentExam(studentExamRepo)

	studentExamController := studentExamInfrastructure.NewStudentExamController(
		createStudentExam,
		getAllStudentExams,
		getStudentExamByID,
		deleteStudentExam,
	)

	resultRepo := resultsInfrastructure.NewResultsRepository(db)

	createResult := resultsApplication.NewCreateResult(resultRepo)
	getAllResultsByExam := resultsApplication.NewGetAllResultsByExam(resultRepo)
	deleteAllResultsByExam := resultsApplication.NewDeleteAllResultsByExam(resultRepo)

	resultController := resultsInfrastructure.NewResultsController(
		createResult,
		getAllResultsByExam,
		deleteAllResultsByExam,
	)

	mux := http.NewServeMux()

	database.RegisterUserRoutes(mux, userController)
	database.RegisterCategoryRoutes(mux, categoryController)
	database.RegisterExamRoutes(mux, examController)
	database.RegisterQuestionRoutes(mux, questionController)
	database.RegisterStudentExamRoutes(mux, studentExamController)
	database.RegisterResultRoutes(mux, resultController)

	handlerWithCORS := enableCORS(mux)

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	log.Println("🚀 Servidor corriendo en http://localhost:" + port)
	log.Fatal(http.ListenAndServe(":"+port, handlerWithCORS))
}

func enableCORS(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Methods", "GET, POST, OPTIONS")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type")

		if r.Method == http.MethodOptions {
			w.WriteHeader(http.StatusOK)
			return
		}

		next.ServeHTTP(w, r)
	})
}
