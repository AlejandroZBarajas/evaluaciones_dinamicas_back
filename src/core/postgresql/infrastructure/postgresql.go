package database

import (
	"database/sql"
	"fmt"
	"os"

	_ "github.com/lib/pq"
)

// DB es la instancia global de la base de datos
var DB *sql.DB

// Connect abre la conexión a la base de datos PostgreSQL
func Connect() error {
	host := os.Getenv("DB_HOST")
	port := os.Getenv("DB_PORT")
	user := os.Getenv("DB_USER")
	password := os.Getenv("DB_PASSWORD")
	dbname := os.Getenv("DB_NAME")

	// Cadena de conexión
	psqlInfo := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)

	var err error
	DB, err = sql.Open("postgres", psqlInfo)
	if err != nil {
		return err
	}

	// Verificamos la conexión
	err = DB.Ping()
	if err != nil {
		return err
	}

	fmt.Println("✅ Conexión a la base de datos establecida")
	return nil
}
