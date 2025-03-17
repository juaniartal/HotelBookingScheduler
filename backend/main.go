package main

import (
	"log"
	"net/http"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
)

func main() {
	ConnectDatabase()

	r := mux.NewRouter()

	// Rutas
	r.HandleFunc("/bookings", CreateBooking).Methods("POST")
	r.HandleFunc("/bookings", GetBookings).Methods("GET")
	r.HandleFunc("/bookings/{id}", GetBookingByID).Methods("GET")
	r.HandleFunc("/bookings/{id}", UpdateBooking).Methods("PUT")
	r.HandleFunc("/bookings/{id}", DeleteBooking).Methods("DELETE")

	corsHandler := handlers.CORS(
		handlers.AllowedOrigins([]string{"http://localhost:5173"}),
		handlers.AllowedMethods([]string{"GET", "POST", "PUT", "DELETE", "OPTIONS"}),
		handlers.AllowedHeaders([]string{"Content-Type", "Authorization"}),
	)

	log.Println("🚀 Servidor corriendo en http://localhost:8080")
	http.ListenAndServe(":8080", corsHandler(r))
}
