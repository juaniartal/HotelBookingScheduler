package main

type Booking struct {
	ID           int      `json:"id"`
	GuestName    string   `json:"guest_name"`
	CheckIn      string   `json:"check_in"`
	CheckOut     string   `json:"check_out"`
	GuestsCount  int      `json:"guests_count"`
	Price        *float64 `json:"price"`
	Department   string   `json:"department"`
	HasVehicle   bool     `json:"has_vehicle"`
	LicensePlate *string  `json:"license_plate,omitempty"`
	CarBrand     *string  `json:"car_brand,omitempty"`
	Is4x4        *bool    `json:"is_4x4,omitempty"`
}
