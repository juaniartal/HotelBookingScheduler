package main

type Booking struct {
	ID           int      `json:"id"`
	GuestName    string   `json:"guest_name"`
	CheckIn      string   `json:"check_in"`
	CheckInTime  string   `json:"check_in_time"`
	CheckOut     string   `json:"check_out"`
	CheckOutTime string   `json:"check_out_time"`
	GuestsCount  int      `json:"guests_count"`
	Price        *float64 `json:"price"`
	Department   string   `json:"department"`
	HasVehicle   bool     `json:"has_vehicle"`
	LicensePlate *string  `json:"license_plate,omitempty"`
	CarBrand     *string  `json:"car_brand,omitempty"`
	Is4x4        *bool    `json:"is_4x4,omitempty"`
}
