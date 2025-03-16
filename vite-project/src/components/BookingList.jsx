import { useState, useEffect } from "react";

export default function BookingList() {
  const [bookings, setBookings] = useState(null); // Estado inicial como null

  useEffect(() => {
    fetch("/api/bookings")
      .then((res) => res.json())
      .then((data) => setBookings(data || [])) // Asegurar que sea un array
      .catch((error) => {
        console.error("Error fetching bookings:", error);
        setBookings([]); // Si hay error, usa un array vacío
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
  
      console.log("Reserva eliminada con éxito");
    } catch (error) {
      console.error("Error al eliminar la reserva", error);
    }
  };
  
  

  if (!bookings) {
    return <p>Cargando reservas...</p>;
  }

  if (bookings.length === 0) {
    return <p>No hay reservas registradas.</p>;
  }

  return (
    <ul className="space-y-2">
      {bookings.map((booking) => (
        <li key={booking.id} className="border p-4 rounded-lg flex justify-between items-center">
          <div>
            <p className="font-bold">{booking.guest_name}</p>
            <p>{booking.check_in} - {booking.check_out}</p>
            <p>{booking.guests_count} personas - ${booking.price}</p>
          </div>
          <div>
            <button className="bg-yellow-500 text-white px-3 py-1 mx-1 rounded">✏️ Editar</button>
            <button onClick={() => handleDelete(booking.id)} className="bg-red-500 text-white px-3 py-1 rounded">🗑️ Eliminar</button>
          </div>
        </li>
      ))}
    </ul>
  );
}