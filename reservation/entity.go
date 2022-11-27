package reservation

import "time"

type Reservation struct {
	ID          string `json:"id"`
	UserID      string `json:"user_id"`
	BookingDate string `json:"booking_date"`
	CreatedAt   time.Time
	UpdatedAt   time.Time
}
