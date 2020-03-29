package passengers

type PassengerInfo struct {
	FirstName        string `json:"firstname"`
	LastName         string `json:"lastname"`
	PassportNumber   int    `json:"passportnumber"`
	DateOfBirth      string `json:"dateofbirth"`
	Email            string `json:"email"`
	CreditcardNumber int    `json:"creditcardnumber"`
}
