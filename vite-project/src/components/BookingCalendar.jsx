import { useState } from "react";
import Calendar from "react-calendar";
import "react-calendar/dist/Calendar.css";
import "../styles/CalendarStyles.css";

function BookingCalendar({ bookings = [], onBack }) { 
  const [selectedDate, setSelectedDate] = useState(new Date());

  const formatDate = (dateString) => new Date(dateString).toLocaleDateString("es-AR");

  const marcarDiasOcupados = ({ date }) => {
    return bookings.some(({ check_in, check_out }) => {
      const checkIn = new Date(check_in);
      const checkOut = new Date(check_out);
      
      checkOut.setDate(checkOut.getDate() + 1); 
  
      return date >= checkIn && date <= checkOut;
    })
      ? "ocupado"
      : null;
  };
  
  const bookingsForDate = bookings?.filter(({ check_in, check_out }) => {
    const checkIn = new Date(check_in);
    const checkOut = new Date(check_out);
    
    checkOut.setDate(checkOut.getDate() + 1); 
  
    return selectedDate >= checkIn && selectedDate <= checkOut;
  }) || [];

  return (
    <div className="container mx-auto p-5">
      <h2 className="text-2xl font-bold mb-4 text-center">📅 Calendario de Reservas</h2>

      <div className="flex flex-col items-center">
        <Calendar
          onChange={setSelectedDate}
          value={selectedDate}
          tileClassName={marcarDiasOcupados}
          className="custom-calendar"
        />
        <p className="mt-2 text-lg font-semibold text-blue-700">
          Seleccionado: {selectedDate.toLocaleDateString("es-AR")}
        </p>
      </div>

      <div className="mt-6">
        <h3 className="text-xl font-semibold text-center">
          Reservas para el {selectedDate.toLocaleDateString("es-AR")}:
        </h3>

        {bookingsForDate.length === 0 ? (
          <p className="text-center text-gray-600 mt-2">No hay reservas para esta fecha.</p>
        ) : (
          <ul className="space-y-3 mt-4">
            {bookingsForDate.map(({ id, guest_name, check_in, check_out, check_in_time, check_out_time, guests_count, price, department, has_vehicle, car_brand, license_plate, is_4x4 }) => {
              const checkInDate = new Date(check_in).toDateString();
              const isCheckInToday = checkInDate === selectedDate.toDateString();

              return (
                <li
                  key={id}
                  className={`border p-4 rounded-lg shadow-lg text-white ${
                    isCheckInToday ? "bg-green-500" : "bg-blue-500"
                  }`}
                >
                  <p className="text-lg font-bold">🛎️ {guest_name}</p>
                  <p>📅 {formatDate(check_in)} {check_in_time && `(${check_in_time})`} - {formatDate(check_out)} {check_out_time && `(${check_out_time})`}</p>
                  <p>👥 {guests_count} personas - 💲{price}</p>
                  <p>🏨 Departamento: {department || "No especificado"}</p>

                  {has_vehicle && (
                    <div className="mt-2 p-2 bg-blue-700 rounded">
                      <p>🚗 Vehículo: {car_brand} ({license_plate}) {is_4x4 && "🛻 4x4"}</p>
                    </div>
                  )}

                  {isCheckInToday && <p className="text-sm font-semibold mt-2">🟢 Entra hoy</p>}
                </li>
              );
            })}
          </ul>
        )}
      </div>

      <div className="flex justify-center mt-6">
        <button onClick={onBack} className="bg-blue-700 text-white p-2 rounded shadow-md hover:bg-blue-900">
          🔙 Volver
        </button>
      </div>
    </div>
  );
}

export default BookingCalendar;
