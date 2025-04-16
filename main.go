package main

import (
	"log"

	database "evaluaciones/src/core/postgresql/infrastructure"

	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("❌ Error al cargar .env:", err)
	}

	err = database.Connect()
	if err != nil {
		log.Fatal("❌ No se pudo conectar a la base de datos:", err)
	}
}
