import { useState, useEffect } from "react";

function BookingForm({ booking, onBookingAdded, onUpdate }) {
  const [guestName, setGuestName] = useState(booking?.guest_name || "");
  const [checkIn, setCheckIn] = useState(booking?.check_in || "");
  const [checkInTime, setCheckInTime] = useState(booking?.check_in_time || "");
  const [checkOut, setCheckOut] = useState(booking?.check_out || "");
  const [checkOutTime, setCheckOutTime] = useState(booking?.check_out_time || "");
  const [guestsCount, setGuestsCount] = useState(booking?.guests_count || 1);
  const [price, setPrice] = useState(booking?.price || "");
  const [department, setDepartment] = useState(booking?.department || "");

  const [hasVehicle, setHasVehicle] = useState(booking?.has_vehicle || false);
  const [licensePlate, setLicensePlate] = useState(booking?.license_plate || "");
  const [carBrand, setCarBrand] = useState(booking?.car_brand || "");
  const [is4x4, setIs4x4] = useState(booking?.is_4x4 || false);

  useEffect(() => {
    if (booking) {
      setGuestName(booking.guest_name);
      setCheckIn(booking.check_in);
      setCheckInTime(booking.check_in_time || "");
      setCheckOut(booking.check_out);
      setCheckOutTime(booking.check_out_time || "");
      setGuestsCount(booking.guests_count);
      setPrice(booking.price);
      setDepartment(booking.department || "");
      setHasVehicle(booking.has_vehicle || false);
      setLicensePlate(booking.license_plate || "");
      setCarBrand(booking.car_brand || "");
      setIs4x4(booking.is_4x4 || false);
    }
  }, [booking]);

  const handleSubmit = async (e) => {
    e.preventDefault();

    // Formatea las horas a HH:MM:SS
    const formattedCheckInTime = checkInTime ? `${checkInTime}:00` : "";
    const formattedCheckOutTime = checkOutTime ? `${checkOutTime}:00` : "";

    const newBooking = {
      id: booking?.id,
      guest_name: guestName,
      check_in: checkIn,
      check_in_time: formattedCheckInTime, // Usamos el formato correcto
      check_out: checkOut,
      check_out_time: formattedCheckOutTime, // Usamos el formato correcto
      guests_count: guestsCount,
      price: price ? parseFloat(price) : 0,
      department: department || null,
      has_vehicle: hasVehicle,
      license_plate: hasVehicle ? licensePlate : null,
      car_brand: hasVehicle ? carBrand : null,
      is_4x4: hasVehicle ? is4x4 : false,
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
        setCheckInTime("");
        setCheckOut("");
        setCheckOutTime("");
        setGuestsCount(1);
        setPrice("");
        setDepartment("");
        setHasVehicle(false);
        setLicensePlate("");
        setCarBrand("");
        setIs4x4(false);
      } else {
        const errorText = await res.text();
        alert(`Error: ${errorText}`);
      }
    }
  };

  return (
    <form onSubmit={handleSubmit} className="my-4 space-y-2">
      <label>Nombre del huésped</label>
      <input type="text" value={guestName} onChange={(e) => setGuestName(e.target.value)} className="border p-2 w-full" required />

      <label>Fecha de Check-In</label>
      <input type="date" value={checkIn} onChange={(e) => setCheckIn(e.target.value)} className="border p-2 w-full" required />

      <label>Hora de Check-In</label>
      <input type="time" value={checkInTime} onChange={(e) => setCheckInTime(e.target.value)} className="border p-2 w-full" />

      <label>Fecha de Check-Out</label>
      <input type="date" value={checkOut} onChange={(e) => setCheckOut(e.target.value)} className="border p-2 w-full" required />

      <label>Hora de Check-Out</label>
      <input type="time" value={checkOutTime} onChange={(e) => setCheckOutTime(e.target.value)} className="border p-2 w-full" />

      <label>Cantidad de Personas</label>
      <input type="number" value={guestsCount} onChange={(e) => setGuestsCount(parseInt(e.target.value))} className="border p-2 w-full" required min="1" />

      <label>Precio ($)</label>
      <input type="number" step="0.01" value={price} onChange={(e) => setPrice(e.target.value)} className="border p-2 w-full" />

      <label>Departamento</label>
      <input type="text" value={department} onChange={(e) => setDepartment(e.target.value)} className="border p-2 w-full" />

      <label>
        ¿Tiene vehículo?
        <input type="checkbox" checked={hasVehicle} onChange={(e) => setHasVehicle(e.target.checked)} className="mr-2" />
      </label>

      {hasVehicle && (
        <>
          <label>Patente del vehículo</label>
          <input type="text" value={licensePlate} onChange={(e) => setLicensePlate(e.target.value)} className="border p-2 w-full" />

          <label>Marca del auto</label>
          <input type="text" value={carBrand} onChange={(e) => setCarBrand(e.target.value)} className="border p-2 w-full" />

          <label>
            ¿Es 4x4?
            <input type="checkbox" checked={is4x4} onChange={(e) => setIs4x4(e.target.checked)} className="mr-2" />
          </label>
        </>
      )}

      <button className="bg-blue-500 text-white p-2 w-full">{booking ? "Actualizar Reserva" : "Agregar Reserva"}</button>
    </form>
  );
}

export default BookingForm;
