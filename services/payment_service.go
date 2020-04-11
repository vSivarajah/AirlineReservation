package services

import (
	payments "github.com/vsivarajah/AirlineReservation/domain/payments"
)

var (
	PaymentService paymentServiceInterface = &paymentService{}
)

type paymentService struct{}

type paymentServiceInterface interface {
	CreatePayment(*payments.Payment) (error, bool)
	GetPayment() (payments.Payments, error)
}

func (s *paymentService) CreatePayment(payment *payments.Payment) (error, bool) {
	return payments.CreatePayment(payment)
}

func (s *paymentService) GetPayment() (payments.Payments, error) {
	return payments.GetPayment()
}
