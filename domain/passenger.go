package domain

type PassengerInfo struct {
	FirstName      string `json:"firstname"`
	LastName       string `json:"lastname"`
	PassportNumber int    `json:"passportnumber"`
	DateOfBirth    string `json:"dateofbirth"`
}
