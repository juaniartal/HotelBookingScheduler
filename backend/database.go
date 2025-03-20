package main

import (
	"database/sql"
	"log"
	"os"

	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

var db *sql.DB

func ConnectDatabase() {
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("❌ No se pudo cargar el archivo .env: %v", err)
	}

	dbUser := os.Getenv("DB_USER")
	dbPassword := os.Getenv("DB_PASSWORD")
	dbName := os.Getenv("DB_NAME")
	connStr := "user=" + dbUser + " password=" + dbPassword + " dbname=" + dbName + " search_path=public sslmode=disable"

	var err2 error
	db, err2 = sql.Open("postgres", connStr)
	if err2 != nil {
		log.Fatalf("❌ Error al conectar a la base de datos: %v", err2)
	}

	err2 = db.Ping()
	if err2 != nil {
		log.Fatalf("❌ No se pudo hacer ping a la base de datos: %v", err2)
	}

	log.Println("✅ Conectado a la base de datos correctamente")
}
