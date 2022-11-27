package reservation

type Reservation struct {
	ID          string `json:"id"`
	UserID      string `json:"user_id"`
	BookingDate string `json:"booking_date"`
}
