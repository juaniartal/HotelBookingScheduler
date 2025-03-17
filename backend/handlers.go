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
	rows, err := db.Query("SELECT id, guest_name, check_in, check_out, guests_count, price, department FROM public.bookings")
	if err != nil {
		log.Println("❌ Error en la consulta a la base de datos:", err)
		http.Error(w, fmt.Sprintf("Database error: %v", err), http.StatusInternalServerError)
		return
	}
	defer rows.Close()

	var bookings []Booking
	for rows.Next() {
		var booking Booking
		var price sql.NullFloat64
		var department sql.NullString

		if err := rows.Scan(&booking.ID, &booking.GuestName, &booking.CheckIn, &booking.CheckOut, &booking.GuestsCount, &price, &department); err != nil {
			log.Println("❌ Error escaneando datos:", err)
			http.Error(w, fmt.Sprintf("Error scanning data: %v", err), http.StatusInternalServerError)
			return
		}

		if price.Valid {
			booking.Price = &price.Float64
		} else {
			booking.Price = nil
		}

		if department.Valid {
			booking.Department = department.String
		} else {
			booking.Department = "No especificado"
		}

		bookings = append(bookings, booking)
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(bookings)
}

func CreateBooking(w http.ResponseWriter, r *http.Request) {
	var booking Booking

	if err := json.NewDecoder(r.Body).Decode(&booking); err != nil {
		log.Println("❌ Error decodificando JSON:", err)
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}

	log.Printf("📥 Reserva recibida: %+v\n", booking)

	if booking.CheckIn == "" || booking.CheckOut == "" {
		http.Error(w, "Missing check-in or check-out date", http.StatusBadRequest)
		return
	}

	if booking.Price == nil {
		defaultPrice := 0.0
		booking.Price = &defaultPrice
	}

	_, err := db.Exec("INSERT INTO bookings (guest_name, check_in, check_out, guests_count, price, department) VALUES ($1, $2, $3, $4, $5, $6)",
		booking.GuestName, booking.CheckIn, booking.CheckOut, booking.GuestsCount, booking.Price, booking.Department)

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
	err := db.QueryRow("SELECT id, guest_name, check_in, check_out, guests_count, price, department FROM public.bookings WHERE id = $1", id).
		Scan(&booking.ID, &booking.GuestName, &booking.CheckIn, &booking.CheckOut, &booking.GuestsCount, &booking.Price, &booking.Department)

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
		SET guest_name=$1, check_in=$2, check_out=$3, guests_count=$4, price=$5, department=$6 
		WHERE id=$7`,
		booking.GuestName, booking.CheckIn, booking.CheckOut, booking.GuestsCount, booking.Price, booking.Department, id)

	if err != nil {
		http.Error(w, "Error updating booking", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(booking)
}
