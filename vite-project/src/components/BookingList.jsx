import { useState, useEffect } from "react"; 
import BookingForm from "./BookingForm";

export default function BookingList() {
  const [bookings, setBookings] = useState(null);
  const [editingBooking, setEditingBooking] = useState(null);

  useEffect(() => {
    fetch("http://localhost:8080/bookings")
      .then((res) => res.json())
      .then((data) => setBookings(data || []))
      .catch((error) => {
        console.error("Error fetching bookings:", error);
        setBookings([]);
      });
  }, []);

  const handleDelete = async (id) => {
    try {
      const response = await fetch(`http://localhost:8080/bookings/${id}`, { method: "DELETE" });

      if (!response.ok) throw new Error("No se pudo eliminar la reserva");

      setBookings((prev) => prev.filter((b) => b.id !== id));
    } catch (error) {
      console.error("Error al eliminar la reserva", error);
    }
  };

  const handleEdit = (booking) => {
    setEditingBooking(booking);
  };

  const handleCancel = () => {
    setEditingBooking(null);
  };

  const handleUpdate = async (updatedBooking) => {
    try {
      const response = await fetch(`http://localhost:8080/bookings/${updatedBooking.id}`, {
        method: "PUT",
        headers: { "Content-Type": "application/json" },
        body: JSON.stringify(updatedBooking),
      });

      if (!response.ok) throw new Error("No se pudo actualizar la reserva");

      setBookings((prev) => prev.map((b) => (b.id === updatedBooking.id ? updatedBooking : b)));
      setEditingBooking(null);
    } catch (error) {
      console.error("Error al actualizar la reserva", error);
    }
  };

  if (!bookings) return <p>Cargando reservas...</p>;
  if (bookings.length === 0) return <p className="no-reservations">No hay reservas registradas.</p>;

  return (
    <div className="max-w-3xl mx-auto mt-6">
      {editingBooking ? (
        <BookingForm booking={editingBooking} onBookingAdded={() => setEditingBooking(null)} onUpdate={handleUpdate} onCancel={handleCancel} />
      ) : (
        <ul className="space-y-4">
          {bookings.map((booking) => (
            <li key={booking.id} className="border p-4 rounded-lg shadow-md bg-white">
              <p className="text-lg font-bold">{booking.guest_name}</p>
              <p className="text-gray-700">
                🏨 {booking.department || "Sin especificar"}
              </p>
              <p className="text-sm text-gray-600">
  📅 {new Date(booking.check_in).toLocaleDateString()} 
  {booking.check_in_time && ` (${new Date(booking.check_in_time).toLocaleTimeString()})`} → 
  {new Date(booking.check_out).toLocaleDateString()} 
  {booking.check_out_time && ` (${new Date(booking.check_out_time).toLocaleTimeString()})`}
</p>

              <p className="text-sm text-gray-600">
                👥 {booking.guests_count} personas - 💰 ${booking.price}
              </p>

              {booking.has_vehicle && (
                <div className="mt-2 p-3 bg-blue-50 border-l-4 border-blue-400 rounded">
                  <p className="text-blue-700 font-semibold">🚗 Vehículo</p>
                  <p className="text-blue-800 text-sm">Marca: {booking.car_brand}</p>
                  <p className="text-blue-800 text-sm">Patente: {booking.license_plate}</p>
                  {booking.is_4x4 && <p className="text-blue-800 text-sm">🛻 Es 4x4</p>}
                </div>
              )}

              <div className="mt-3 flex gap-2">
                <button onClick={() => handleEdit(booking)} className="bg-yellow-400 hover:bg-yellow-500 text-white px-3 py-1 rounded">
                  ✏️ Editar
                </button>
                <button onClick={() => handleDelete(booking.id)} className="bg-red-500 hover:bg-red-600 text-white px-3 py-1 rounded">
                  🗑️ Eliminar
                </button>
              </div>
            </li>
          ))}
        </ul>
      )}
    </div>
  );
}
