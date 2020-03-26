package domain

type Reservation struct {
	Id         int64         `json:"id"`
	Passenger  PassengerInfo `json:"passenger"`
	FlightInfo Flight        `json:flightinfo`
}
type Reservations []*Reservation

var reservationDetails = []*Reservation{
	&Reservation{
		Id: 1,
		Passenger: PassengerInfo{
			FirstName:      "Vignesh",
			LastName:       "Sivarajah",
			PassportNumber: 1234567,
			DateOfBirth:    "201299",
		},
		FlightInfo: Flight{
			FlightNumber:      "BOEING777",
			OperatingAirlines: "Emirates",
			SourceAirport:     "Oslo",
			TargetAirport:     "Cancun",
		},
	},
}

func CreateFlightDetails(reservation *Reservation) {
	reservationDetails = append(reservationDetails, reservation)
}
func GetReservation() Reservations {
	return reservationDetails
}
