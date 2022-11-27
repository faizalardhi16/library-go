package reservation

import "github.com/google/uuid"

type Service interface {
	CreateReservation(input ReservationInput) (Reservation, error)
}

type service struct {
	repository Repository
}

func NewReservationService(repository Repository) *service {
	return &service{repository}
}

func (s *service) CreateReservation(input ReservationInput) (Reservation, error) {
	inputValue := Reservation{}
	inputValue.ID = uuid.New().String()
	inputValue.UserID = input.UserID
	inputValue.BookingDate = input.BookingDate

	res, err := s.repository.Save(inputValue)

	if err != nil {
		return res, err
	}

	return res, nil
}
