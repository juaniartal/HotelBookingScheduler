package main

import (
	"log"
	"os"
	"testing"
)

// Función para configurar la base de datos para las pruebas
func setupTestDB() {
	// Cambia la variable de entorno a "test" solo para el entorno de pruebas
	err := os.Setenv("DB_ENV", "test")
	if err != nil {
		log.Fatalf("❌ Error al configurar el entorno de test: %v", err)
	}

	// Conecta a la base de datos de prueba
	ConnectDatabase()

	// Crear las tablas necesarias para los tests si es necesario
	_, err = db.Exec(`
		CREATE TABLE IF NOT EXISTS bookings (
			id SERIAL PRIMARY KEY,
			guest_name TEXT,
			check_in TEXT,
			check_out TEXT,
			guests_count INT,
			price REAL,
			department TEXT,
			has_vehicle BOOLEAN,
			license_plate TEXT,
			car_brand TEXT,
			is_4x4 BOOLEAN,
			deleted BOOLEAN
		)
	`)
	if err != nil {
		log.Fatalf("❌ Error al crear tablas en la base de datos de prueba: %v", err)
	}
}

// Función para limpiar la base de datos después de las pruebas
func teardownTestDB() {
	// Aquí puedes eliminar los datos si lo deseas
	_, err := db.Exec("DROP TABLE IF EXISTS bookings")
	if err != nil {
		log.Fatalf("❌ Error al limpiar la base de datos de prueba: %v", err)
	}

	// Cierra la conexión
	db.Close()
}

// Prueba de ejemplo
func TestCreateBooking(t *testing.T) {
	setupTestDB()
	defer teardownTestDB()

	// Insertar una reserva
	_, err := db.Exec("INSERT INTO bookings (guest_name, check_in, check_out, guests_count, price, department, has_vehicle, license_plate, car_brand, is_4x4, deleted) VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11)", "Juan Pérez", "2024-01-01", "2024-01-02", 2, 100.0, "Apt. 101", false, nil, nil, false, false)
	if err != nil {
		t.Fatalf("❌ Error al insertar la reserva: %v", err)
	}

	// Verificar la inserción
	var count int
	err = db.QueryRow("SELECT COUNT(*) FROM bookings WHERE guest_name = $1", "Juan Pérez").Scan(&count)
	if err != nil || count == 0 {
		t.Fatal("❌ La reserva no se guardó correctamente")
	}
}
