import { useState, useEffect } from "react";
import BookingForm from "./BookingForm"; // Importamos el formulario

export default function BookingList() {
  const [bookings, setBookings] = useState(null);
  const [editingBooking, setEditingBooking] = useState(null); // Estado para edición

  useEffect(() => {
    fetch("/api/bookings")
      .then((res) => res.json())
      .then((data) => setBookings(data || []))
      .catch((error) => {
        console.error("Error fetching bookings:", error);
        setBookings([]);
      });
  }, []);

  const handleDelete = async (id) => {
    try {
      const response = await fetch(`http://localhost:8080/bookings/${id}`, {
        method: "DELETE",
      });

      if (!response.ok) {
        throw new Error("No se pudo eliminar la reserva");
      }

      setBookings((prev) => prev.filter((b) => b.id !== id)); // Eliminamos de la UI
    } catch (error) {
      console.error("Error al eliminar la reserva", error);
    }
  };

  const handleEdit = (booking) => {
    setEditingBooking(booking); // Cargar la reserva en edición
  };

  const handleUpdate = async (updatedBooking) => {
    try {
      const response = await fetch(`http://localhost:8080/bookings/${updatedBooking.id}`, {
        method: "PUT",
        headers: { "Content-Type": "application/json" },
        body: JSON.stringify(updatedBooking),
      });

      if (!response.ok) {
        throw new Error("No se pudo actualizar la reserva");
      }

      setBookings((prev) => prev.map((b) => (b.id === updatedBooking.id ? updatedBooking : b)));
      setEditingBooking(null); // Salir del modo edición
    } catch (error) {
      console.error("Error al actualizar la reserva", error);
    }
  };

  if (!bookings) return <p>Cargando reservas...</p>;
  if (bookings.length === 0) return <p>No hay reservas registradas.</p>;

  return (
    <div>
      {editingBooking && (
        <BookingForm
          booking={editingBooking}
          onBookingAdded={() => setEditingBooking(null)} // Ocultar el formulario
          onUpdate={handleUpdate} // Pasamos la función de actualización
        />
      )}

      <ul className="space-y-2">
        {bookings.map((booking) => (
          <li key={booking.id} className="border p-4 rounded-lg flex justify-between items-center">
            <div>
              <p className="font-bold">{booking.guest_name}</p>
              <p>{booking.check_in} - {booking.check_out}</p>
              <p>{booking.guests_count} personas - ${booking.price}</p>
            </div>
            <div>
              <button
                onClick={() => handleEdit(booking)}
                className="bg-yellow-500 text-white px-3 py-1 mx-1 rounded"
              >
                ✏️ Editar
              </button>
              <button
                onClick={() => handleDelete(booking.id)}
                className="bg-red-500 text-white px-3 py-1 rounded"
              >
                🗑️ Eliminar
              </button>
            </div>
          </li>
        ))}
      </ul>
    </div>
  );
}
