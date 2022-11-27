package transaction

import "gorm.io/gorm"

type Repository interface {
	Save(transaction Transaction) (Transaction, error)
	FindTransaction() ([]Transaction, error)
}

type repository struct {
	db *gorm.DB
}

func NewTransactionRepository(db *gorm.DB) *repository {
	return &repository{db}
}

func (r *repository) Save(transaction Transaction) (Transaction, error) {
	err := r.db.Create(&transaction).Error

	if err != nil {
		return transaction, err
	}

	return transaction, nil
}

func (r *repository) FindTransaction() ([]Transaction, error) {
	var transaction []Transaction

	err := r.db.Find(&transaction).Error

	if err != nil {
		return transaction, err
	}

	return transaction, nil
}
