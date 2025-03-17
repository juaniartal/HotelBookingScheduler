// BookingCalendar.jsx
import { useState } from "react";
import Calendar from "react-calendar";
import "react-calendar/dist/Calendar.css";
import "../styles/CalendarStyles.css";

function BookingCalendar({ bookings, onBack }) {
  const [selectedDate, setSelectedDate] = useState(new Date());

  const formatDate = (dateString) => new Date(dateString).toLocaleDateString("es-AR");

  const marcarDiasOcupados = ({ date }) => {
    return bookings.some(({ check_in, check_out }) => {
      const checkIn = new Date(check_in);
      const checkOut = new Date(check_out);
      return date >= checkIn && date <= checkOut;
    }) ? "ocupado" : null;
  };

  const bookingsForDate = bookings.filter(({ check_in, check_out }) => {
    const checkIn = new Date(check_in);
    const checkOut = new Date(check_out);
    return selectedDate >= checkIn && selectedDate <= checkOut;
  });

  return (
    <div className="container mx-auto p-5">
      <h2 className="text-2xl font-bold mb-4 text-center">📅 Calendario de Reservas</h2>

      <div className="flex justify-center mb-4">
        <Calendar 
          onChange={setSelectedDate} 
          value={selectedDate} 
          tileClassName={marcarDiasOcupados} 
          className="custom-calendar"
        />
      </div>

      <div className="mt-4">
        <h3 className="text-xl font-semibold text-center">
          Reservas para el {selectedDate.toLocaleDateString("es-AR")}:
        </h3>

        {bookingsForDate.length === 0 ? (
          <p className="text-center text-gray-600">No hay reservas para esta fecha.</p>
        ) : (
          <ul className="space-y-2">
            {bookingsForDate.map(({ id, guest_name, check_in, check_out, guests_count, price, department }) => (
              <li key={id} className="border p-4 rounded-lg shadow-lg bg-white">
                <p className="font-bold">{guest_name}</p>
                <p>📅 {formatDate(check_in)} - {formatDate(check_out)}</p>
                <p>👥 {guests_count} personas - 💲{price}</p>
                <p>🏨 Departamento: {department || "No especificado"}</p>
              </li>
            ))}
          </ul>
        )}
      </div>

      <div className="flex justify-center mt-4">
        <button 
          onClick={onBack} 
          className="bg-blue-500 text-white p-2 rounded shadow-md hover:bg-blue-700"
        >
          🔙 Volver
        </button>
      </div>
    </div>
  );
}

export default BookingCalendar;