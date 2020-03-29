package domain

type Payment struct {
	PaymentId int `json:"paymentid"`
	Reservation
}
