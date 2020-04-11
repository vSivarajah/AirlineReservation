package services

import (
	"github.com/vsivarajah/AirlineReservation/domain/reservations"
)

var (
	ReservationService reservationServiceInterface = &reservationService{}
)

type reservationService struct{}

type reservationServiceInterface interface {
	CreateReservation(*reservations.Reservation) bool
	GetReservationDetails() reservations.Reservations
	FindReservationById(int) (*reservations.Reservation, int, error)
	UpdateReservation(int, *reservations.Reservation) error
	DeleteReservation(int) int
}

func (s *reservationService) CreateReservation(reservation *reservations.Reservation) bool {
	return reservations.CreateReservation(reservation)
}

func (s *reservationService) GetReservationDetails() reservations.Reservations {
	return reservations.GetReservationDetails()
}

func (s *reservationService) FindReservationById(id int) (*reservations.Reservation, int, error) {
	return reservations.FindReservationById(id)
}

func (s *reservationService) UpdateReservation(id int, r *reservations.Reservation) error {
	return reservations.UpdateReservation(id, r)
}

func (s *reservationService) DeleteReservation(id int) int {
	return reservations.DeleteReservation(id)
}
