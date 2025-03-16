package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func GetBookings(w http.ResponseWriter, r *http.Request) {
	rows, err := db.Query("SELECT id, guest_name, check_in, check_out, guests_count, price FROM public.bookings")
	if err != nil {
		log.Println("❌ Error en la consulta a la base de datos:", err)
		http.Error(w, fmt.Sprintf("Database error: %v", err), http.StatusInternalServerError)
		return
	}
	defer rows.Close()

	var bookings []Booking
	for rows.Next() {
		var booking Booking
		var price sql.NullFloat64 // ✅ Manejo de valores NULL

		if err := rows.Scan(&booking.ID, &booking.GuestName, &booking.CheckIn, &booking.CheckOut, &booking.GuestsCount, &price); err != nil {
			log.Println("❌ Error escaneando datos:", err)
			http.Error(w, fmt.Sprintf("Error scanning data: %v", err), http.StatusInternalServerError)
			return
		}

		if price.Valid {
			booking.Price = &price.Float64
		} else {
			booking.Price = nil
		}

		bookings = append(bookings, booking)
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(bookings)
}

func CreateBooking(w http.ResponseWriter, r *http.Request) {
	var booking Booking

	// Decodificar JSON
	if err := json.NewDecoder(r.Body).Decode(&booking); err != nil {
		log.Println("❌ Error decodificando JSON:", err)
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}

	// 📌 Ver los datos que llegan al backend
	log.Printf("📥 Reserva recibida: %+v\n", booking)

	// Validar fechas
	if booking.CheckIn == "" || booking.CheckOut == "" {
		http.Error(w, "Missing check-in or check-out date", http.StatusBadRequest)
		return
	}

	if booking.Price == nil {
		defaultPrice := 0.0
		booking.Price = &defaultPrice
	}

	// Verificar disponibilidad de fechas
	var count int
	err := db.QueryRow("SELECT COUNT(*) FROM bookings WHERE (check_in, check_out) OVERLAPS ($1, $2)",
		booking.CheckIn, booking.CheckOut).Scan(&count)
	if err != nil {
		log.Println("❌ Error al verificar disponibilidad:", err)
		http.Error(w, "Error checking availability", http.StatusInternalServerError)
		return
	}
	if count > 0 {
		http.Error(w, "Date not available", http.StatusConflict)
		return
	}

	// Insertar la reserva en la base de datos
	_, err = db.Exec("INSERT INTO bookings (guest_name, check_in, check_out, guests_count, price) VALUES ($1, $2, $3, $4, $5)",
		booking.GuestName, booking.CheckIn, booking.CheckOut, booking.GuestsCount, booking.Price)

	if err != nil {
		log.Println("❌ Error al insertar la reserva:", err)
		http.Error(w, "Could not create booking", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(booking)
}
func GetBookingByID(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]

	var booking Booking
	err := db.QueryRow("SELECT id, guest_name, check_in, check_out, guests_count, price FROM public.bookings WHERE id = $1", id).
		Scan(&booking.ID, &booking.GuestName, &booking.CheckIn, &booking.CheckOut, &booking.GuestsCount, &booking.Price)

	if err != nil {
		http.Error(w, "Booking not found", http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(booking)
}

func DeleteBooking(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]

	_, err := db.Exec("DELETE FROM public.bookings WHERE id = $1", id)
	if err != nil {
		http.Error(w, "Error deleting booking", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}

func UpdateBooking(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]

	var booking Booking
	if err := json.NewDecoder(r.Body).Decode(&booking); err != nil {
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}

	_, err := db.Exec(`UPDATE public.bookings 
		SET guest_name=$1, check_in=$2, check_out=$3, guests_count=$4, price=$5 
		WHERE id=$6`,
		booking.GuestName, booking.CheckIn, booking.CheckOut, booking.GuestsCount, booking.Price, id)

	if err != nil {
		http.Error(w, "Error updating booking", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(booking)
}
