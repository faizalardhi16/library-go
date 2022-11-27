package reservation

import "gorm.io/gorm"

type Repository interface {
	Save(reservation Reservation) (Reservation, error)
}

type repository struct {
	db *gorm.DB
}

func NewReservationRepository(db *gorm.DB) *repository {
	return &repository{db}
}

func (r *repository) Save(reservation Reservation) (Reservation, error) {
	err := r.db.Create(&reservation).Error

	if err != nil {
		return reservation, err
	}

	return reservation, nil
}
