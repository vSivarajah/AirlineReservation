package services

var (
	PaymentService paymentServiceInterface = &paymentService{}
)

type paymentService struct{}

type paymentServiceInterface interface {
	CreatePayment(*payments.Payment)
	GetPayment() payments.Payments
}

func (s *paymentService) CreatePayment(payment *payments.Payment) {
	payments.CreatePayment(payment)
}

func (s *paymentService) GetPayment() payments.Payments {
	return payments.GetPayment()
}
