import { useState } from "react";

function BookingForm({ onBookingAdded }) {
  const [guestName, setGuestName] = useState("");
  const [checkIn, setCheckIn] = useState("");
  const [checkOut, setCheckOut] = useState("");
  const [guestsCount, setGuestsCount] = useState(1);
  const [price, setPrice] = useState("");

  const handleSubmit = async (e) => {
    e.preventDefault();

    const newBooking = {
      guest_name: guestName,
      check_in: checkIn,
      check_out: checkOut,
      guests_count: guestsCount,
      price: price ? parseFloat(price) : 0, // 👈 Si está vacío, enviamos 0
    };

    console.log("📤 Enviando reserva:", newBooking); // 👀 Ver qué se envía al backend

    const res = await fetch("http://localhost:8080/bookings", { // 👈 URL correcta
      method: "POST",
      headers: { "Content-Type": "application/json" },
      body: JSON.stringify(newBooking),
    });

    if (res.ok) {
      onBookingAdded();
      setGuestName("");
      setCheckIn("");
      setCheckOut("");
      setGuestsCount(1);
      setPrice("");
    } else {
      const errorText = await res.text();
      console.error("❌ Error al agregar reserva:", errorText);
      alert(`Error: ${errorText}`); // 👈 Ver error exacto
    }
  };

  return (
    <form onSubmit={handleSubmit} className="my-4 space-y-2">
      <label>Nombre del huésped</label>
      <input
        type="text"
        value={guestName}
        onChange={(e) => setGuestName(e.target.value)}
        className="border p-2 w-full"
        required
      />

      <label>Fecha de Check-In</label>
      <input
        type="date"
        value={checkIn}
        onChange={(e) => setCheckIn(e.target.value)}
        className="border p-2 w-full"
        required
      />

      <label>Fecha de Check-Out</label>
      <input
        type="date"
        value={checkOut}
        onChange={(e) => setCheckOut(e.target.value)}
        className="border p-2 w-full"
        required
      />

      <label>Cantidad de Personas</label>
      <input
        type="number"
        value={guestsCount}
        onChange={(e) => setGuestsCount(parseInt(e.target.value))}
        className="border p-2 w-full"
        required
        min="1"
      />

      <label>Precio ($)</label>
      <input
        type="number"
        step="0.01"
        value={price}
        onChange={(e) => setPrice(e.target.value)}
        className="border p-2 w-full"
      />

      <button className="bg-blue-500 text-white p-2 w-full">
        Agregar Reserva
      </button>
    </form>
  );
}

export default BookingForm;
