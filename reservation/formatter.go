package reservation

type ResponseReservation struct {
	ID          string `json:"id"`
	UserID      string `json:"user_id"`
	BookingDate string `json:"booking_date"`
}

func ResponseReservationFormat(reservation Reservation) ResponseReservation {
	f := ResponseReservation{}

	f.UserID = reservation.UserID
	f.BookingDate = reservation.BookingDate
	f.ID = reservation.ID

	return f
}
