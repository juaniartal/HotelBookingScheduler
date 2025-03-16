package main

type Booking struct {
	ID          int      `json:"id"`
	GuestName   string   `json:"guest_name"`
	CheckIn     string   `json:"check_in"`
	CheckOut    string   `json:"check_out"`
	GuestsCount int      `json:"guests_count"`
	Price       *float64 `json:"price"` // ✅ acepta valores NULL
}
