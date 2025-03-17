package main

import (
	"database/sql"
	"log"

	_ "github.com/lib/pq"
)

var db *sql.DB

func ConnectDatabase() {
	var err error
	connStr := "user=postgres password=JuanDomingoPeron dbname=renata search_path=public sslmode=disable"
	db, err = sql.Open("postgres", connStr)
	if err != nil {
		log.Fatalf("❌ Error al conectar a la base de datos: %v", err)
	}

	err = db.Ping()
	if err != nil {
		log.Fatalf("❌ No se pudo hacer ping a la base de datos: %v", err)
	}

	log.Println("✅ Conectado a la base de datos correctamente")
}
