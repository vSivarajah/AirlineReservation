package domain

import (
	"log"

	"github.com/vsivarajah/AirlineReservation/domain/reservations"
)

type Payment struct {
	PaymentID int `json:"paymentid"`
}

type Payments []*Payment

var PaymentList = []*Payment{}

func CreatePayment(payment *Payment) (error, bool) {
	_, i, _ := FindPaymentById(payment.PaymentID)
	if i != -1 {
		log.Println("Payment has already been made")
		return nil, false
	}

	_, j, _ := reservations.FindReservationById(payment.PaymentID)
	if j == -1 {
		log.Println("Reservation does not exist, it should exist a valid reservation in order to make a payment")
		return nil, false
	}
	PaymentList = append(PaymentList, payment)
	return nil, true
}

func GetPayment() Payments {
	return PaymentList
}

func FindPaymentById(id int) (*Payment, int, error) {
	for i, v := range PaymentList {
		if v.PaymentID == id {
			//Found id
			return v, i, nil
		}
	}
	return nil, -1, nil
}
