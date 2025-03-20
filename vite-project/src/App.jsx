import { useState, useEffect } from "react";
import BookingForm from "./components/BookingForm";
import BookingList from "./components/BookingList";
import BookingCalendar from "./components/BookingCalendar"; 
import "./styles/CalendarStyles.css";

function App() {
  const [bookings, setBookings] = useState([]);
  const [view, setView] = useState("list"); 

  const fetchBookings = () => {
    fetch("http://localhost:8080/bookings")
      .then((res) => res.json())
      .then((data) => setBookings(data))
      .catch((err) => console.error("Error fetching bookings:", err));
  };

  useEffect(() => {
    fetchBookings();
  }, []);

  return (
    <div className="p-5">
      {view === "list" ? (
        <>
          <h1 className="text-2xl font-bold mb-4">Reservas</h1>
          <BookingForm onBookingAdded={fetchBookings} />
          <BookingList bookings={bookings} />
          <button 
            className="mt-4 bg-green-500 text-white p-2 rounded"
            onClick={() => setView("calendar")}
          >
            📅 Ver Calendario
          </button>
        </>
      ) : (
        <BookingCalendar bookings={bookings} onBack={() => setView("list")} />
      )}
    </div>
  );
}

export default App;
