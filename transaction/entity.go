package transaction

import "time"

type Transaction struct {
	ID            string `json:"id"`
	UserID        string `json:"user_id"`
	BookID        string `json:"book_id"`
	ReservationID string `json:"reservation_id"`
	CreatedAt     time.Time
	UpdatedAt     time.Time
}
