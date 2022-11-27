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

func ResponseReservationFormats(reservation []Reservation) []ResponseReservation {

	var f []ResponseReservation

	for _, r := range reservation {
		j := ResponseReservationFormat(r)
		f = append(f, j)
	}

	return f
}
