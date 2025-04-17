package main

import (
	"log"
	"net/http"
	"os"

	database "evaluaciones/src/core/postgresql/infrastructure"
	userApplication "evaluaciones/src/users/application"
	userInfrastructure "evaluaciones/src/users/infrastructure"

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

	mux := http.NewServeMux()
	database.RegisterUserRoutes(mux, userController)
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
