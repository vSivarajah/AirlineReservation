package domain

type Reservation struct {
	Id         int64         `json:"id"`
	Passenger  PassengerInfo `json:"passenger"`
	FlightInfo Flight        `json:flightinfo`
}
type Reservations []*Reservation

var reservationDetails = []*Reservation{}

func CreateFlightDetails(reservation *Reservation) {
	reservationDetails = append(reservationDetails, reservation)
}

func GetReservationDetails() Reservations {
	return reservationDetails
}
