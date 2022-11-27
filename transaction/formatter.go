package transaction

type ResponseTransactionFormatter struct {
	ID            string `json:"id"`
	BookID        string `json:"book_id"`
	UserID        string `json:"user_id"`
	ReservationID string `json:"reservation_id"`
}

func ResponseTransactionFormat(transaction Transaction) ResponseTransactionFormatter {
	f := ResponseTransactionFormatter{}

	f.BookID = transaction.BookID
	f.UserID = transaction.UserID
	f.ID = transaction.ID

	return f
}

func TransactionAllFormat(transaction []Transaction) []ResponseTransactionFormatter {
	transactionFormat := []ResponseTransactionFormatter{}

	for _, r := range transaction {
		transFormat := ResponseTransactionFormat(r)
		transactionFormat = append(transactionFormat, transFormat)
	}

	return transactionFormat
}
