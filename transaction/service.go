package transaction

import "github.com/google/uuid"

type Service interface {
	CreateTransaction(input CreateTransactionInput) (Transaction, error)
	GetTransaction() ([]Transaction, error)
}

type service struct {
	repository Repository
}

func NewTransactionService(repository Repository) *service {
	return &service{repository}
}

func (s *service) CreateTransaction(input CreateTransactionInput) (Transaction, error) {
	transaction := Transaction{}
	transaction.UserID = input.UserID
	transaction.BookID = input.BookID
	transaction.ID = uuid.New().String()

	newTrans, err := s.repository.Save(transaction)

	if err != nil {
		return newTrans, err
	}

	return newTrans, nil
}

func (s *service) GetTransaction() ([]Transaction, error) {
	trans, err := s.repository.FindTransaction()

	if err != nil {
		return trans, err
	}

	return trans, nil
}
