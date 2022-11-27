package transaction

type CreateTransactionInput struct {
	BookID string `json:"book_id"`
	UserID string `json:"user_id"`
}

type CreateInputValue struct {
	BookID string `json:"book_id"`
}

type UpdateTransactionInput struct {
	ReservationID string `json:"reservation_id"`
}
