package main

import (
	"log"
	"net/http"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
)

func main() {
	ConnectDatabase()

	// Crear router
	r := mux.NewRouter()

	// Rutas
	r.HandleFunc("/bookings", CreateBooking).Methods("POST")
	r.HandleFunc("/bookings", GetBookings).Methods("GET")

	// **Habilitar CORS**
	corsHandler := handlers.CORS(
		handlers.AllowedOrigins([]string{"http://localhost:5173"}), // Permitir frontend
		handlers.AllowedMethods([]string{"GET", "POST", "PUT", "DELETE", "OPTIONS"}),
		handlers.AllowedHeaders([]string{"Content-Type", "Authorization"}),
	)

	// Iniciar servidor con CORS
	log.Println("🚀 Servidor corriendo en http://localhost:8080")
	http.ListenAndServe(":8080", corsHandler(r))
}
