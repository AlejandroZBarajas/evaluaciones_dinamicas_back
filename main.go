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

	mux := http.NewServeMux()

	database.RegisterUserRoutes(mux, userController)
	database.RegisterCategoryRoutes(mux, categoryController)
	database.RegisterExamRoutes(mux, examController)

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
