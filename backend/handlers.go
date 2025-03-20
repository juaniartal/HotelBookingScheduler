package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/gorilla/mux"
)

func GetBookings(w http.ResponseWriter, r *http.Request) {
	rows, err := db.Query("SELECT id, guest_name, check_in, check_in_time, check_out, check_out_time, guests_count, price, department FROM public.bookings WHERE deleted = FALSE")
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

		if err := rows.Scan(&booking.ID, &booking.GuestName, &booking.CheckIn, &booking.CheckInTime, &booking.CheckOut, &booking.CheckOutTime, &booking.GuestsCount, &price, &department); err != nil {
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

	// Validar fechas y horas
	if booking.CheckIn == "" || booking.CheckOut == "" || booking.CheckInTime == "" || booking.CheckOutTime == "" {
		http.Error(w, "Missing check-in, check-out date or time", http.StatusBadRequest)
		return
	}

	// Convertir fechas y horas a time.Time
	checkInTime, err := time.Parse("2006-01-02", booking.CheckIn)
	if err != nil {
		http.Error(w, "Invalid check-in date format", http.StatusBadRequest)
		return
	}

	checkOutTime, err := time.Parse("2006-01-02", booking.CheckOut)
	if err != nil {
		http.Error(w, "Invalid check-out date format", http.StatusBadRequest)
		return
	}

	checkInTimeParsed, err := time.Parse("15:04:05", booking.CheckInTime)
	if err != nil {
		http.Error(w, "Invalid check-in time format", http.StatusBadRequest)
		return
	}

	checkOutTimeParsed, err := time.Parse("15:04:05", booking.CheckOutTime)
	if err != nil {
		http.Error(w, "Invalid check-out time format", http.StatusBadRequest)
		return
	}

	// Si el precio es nulo, asignar un precio por defecto
	if booking.Price == nil {
		defaultPrice := 0.0
		booking.Price = &defaultPrice
	}

	_, err = db.Exec(
		"INSERT INTO bookings (guest_name, check_in, check_in_time, check_out, check_out_time, guests_count, price, department, has_vehicle, license_plate, car_brand, is_4x4, deleted) VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12, FALSE)",
		booking.GuestName, checkInTime, checkInTimeParsed, checkOutTime, checkOutTimeParsed, booking.GuestsCount, booking.Price, booking.Department,
		booking.HasVehicle, booking.LicensePlate, booking.CarBrand, booking.Is4x4,
	)

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
	err := db.QueryRow("SELECT id, guest_name, check_in, check_out, guests_count, price, department FROM public.bookings WHERE id = $1 AND deleted = FALSE", id).
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

	_, err := db.Exec("UPDATE public.bookings SET deleted = TRUE WHERE id = $1", id)
	if err != nil {
		http.Error(w, "Error al eliminar la reserva", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
}

func UpdateBooking(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]

	var booking Booking
	if err := json.NewDecoder(r.Body).Decode(&booking); err != nil {
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}

	// Validar fechas y horas
	checkInTime, err := time.Parse("2006-01-02", booking.CheckIn)
	if err != nil {
		http.Error(w, "Invalid check-in date format", http.StatusBadRequest)
		return
	}

	checkOutTime, err := time.Parse("2006-01-02", booking.CheckOut)
	if err != nil {
		http.Error(w, "Invalid check-out date format", http.StatusBadRequest)
		return
	}

	checkInTimeParsed, err := time.Parse("15:04:05", booking.CheckInTime)
	if err != nil {
		http.Error(w, "Invalid check-in time format", http.StatusBadRequest)
		return
	}

	checkOutTimeParsed, err := time.Parse("15:04:05", booking.CheckOutTime)
	if err != nil {
		http.Error(w, "Invalid check-out time format", http.StatusBadRequest)
		return
	}

	_, err = db.Exec(
		"UPDATE bookings SET guest_name=$1, check_in=$2, check_in_time=$3, check_out=$4, check_out_time=$5, guests_count=$6, price=$7, department=$8, has_vehicle=$9, license_plate=$10, car_brand=$11, is_4x4=$12 WHERE id=$13 AND deleted = FALSE",
		booking.GuestName, checkInTime, checkInTimeParsed, checkOutTime, checkOutTimeParsed, booking.GuestsCount, booking.Price, booking.Department,
		booking.HasVehicle, booking.LicensePlate, booking.CarBrand, booking.Is4x4, id,
	)

	if err != nil {
		http.Error(w, "Error updating booking", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(booking)
}
