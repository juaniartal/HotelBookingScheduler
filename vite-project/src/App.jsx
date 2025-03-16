import { useState, useEffect } from "react";
import BookingForm from "./components/BookingForm";
import BookingList from "./components/BookingList";

function App() {
  const [bookings, setBookings] = useState([]);

  const fetchBookings = () => {
    fetch("/api/bookings")
      .then((res) => res.json())
      .then((data) => setBookings(data))
      .catch((err) => console.error("Error fetching bookings:", err));
  };

  useEffect(() => {
    fetchBookings();
  }, []);

  return (
    <div className="p-5">
      <h1 className="text-2xl font-bold">Reservas</h1>
      <BookingForm onBookingAdded={fetchBookings} />
      <BookingList bookings={bookings} />
    </div>
  );
}

export default App;
