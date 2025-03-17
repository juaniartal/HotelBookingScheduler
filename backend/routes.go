package main

import (
	"github.com/gorilla/mux"
)

func SetupRoutes() *mux.Router {
	r := mux.NewRouter()

	r.HandleFunc("/bookings", GetBookings).Methods("GET")
	r.HandleFunc("/bookings/{id}", GetBookingByID).Methods("GET")
	r.HandleFunc("/bookings", CreateBooking).Methods("POST")
	r.HandleFunc("/bookings/{id}", UpdateBooking).Methods("PUT")
	r.HandleFunc("/bookings/{id}", DeleteBooking).Methods("DELETE")

	return r
}
