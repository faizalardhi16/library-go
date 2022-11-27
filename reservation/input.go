package reservation

type ReservationInput struct {
	UserID      string `json:"user_id"`
	BookingDate string `json:"booking_date"`
}
