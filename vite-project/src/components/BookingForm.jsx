import { useState, useEffect } from "react";

function BookingForm({ booking, onBookingAdded, onUpdate }) {
  const [guestName, setGuestName] = useState(booking?.guest_name || "");
  const [checkIn, setCheckIn] = useState(booking?.check_in || "");
  const [checkOut, setCheckOut] = useState(booking?.check_out || "");
  const [guestsCount, setGuestsCount] = useState(booking?.guests_count || 1);
  const [price, setPrice] = useState(booking?.price || "");

  useEffect(() => {
    if (booking) {
      setGuestName(booking.guest_name);
      setCheckIn(booking.check_in);
      setCheckOut(booking.check_out);
      setGuestsCount(booking.guests_count);
      setPrice(booking.price);
    }
  }, [booking]);

  const handleSubmit = async (e) => {
    e.preventDefault();

    const newBooking = {
      id: booking?.id, // Si existe, es edición
      guest_name: guestName,
      check_in: checkIn,
      check_out: checkOut,
      guests_count: guestsCount,
      price: price ? parseFloat(price) : 0,
    };

    if (booking) {
      await onUpdate(newBooking);
    } else {
      const res = await fetch("http://localhost:8080/bookings", {
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
        alert(`Error: ${errorText}`);
      }
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
        {booking ? "Actualizar Reserva" : "Agregar Reserva"}
      </button>
    </form>
  );
}

export default BookingForm;
