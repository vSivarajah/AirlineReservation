package domain

type Payment struct {
	PaymentID int `json:"paymentid"`
}

type Payments []*Payment

var PaymentList = []*Payment{}

func CreatePayment(payment *Payment) {
	PaymentList = append(PaymentList, payment)
}

func GetPayment() Payments {
	return PaymentList
}
