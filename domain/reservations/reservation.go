package reservations

import (
	"log"

	"github.com/vsivarajah/AirlineReservation/domain/flights"
	"github.com/vsivarajah/AirlineReservation/domain/passengers"
)

type Reservation struct {
	Id         int                      `json:"id"`
	IsValid    bool                     `json:"paymentsuccessful"`
	Passenger  passengers.PassengerInfo `json:"passenger"`
	FlightInfo flights.Flight           `json:"flightinfo"`
}
type Reservations []*Reservation

var reservationDetails = []*Reservation{}

// CreateFlightDetails - creates reservation
func CreateFlightDetails(reservation *Reservation) bool {

	_, i, _ := FindReservationById(reservation.Id)

	if i != -1 {
		log.Println("Reservation id exists")
		return false
	} else {
		reservationDetails = append(reservationDetails, reservation)
		return true
	}
}

func GetReservationDetails() Reservations {
	return reservationDetails
}

func FindReservationById(id int) (*Reservation, int, error) {
	for i, v := range reservationDetails {
		if v.Id == id {
			//Found id
			return v, i, nil
		}
	}
	return nil, -1, nil
}

func UpdateReservation(id int, r *Reservation) error {
	value, pos, err := FindReservationById(id)
	if err != nil {
		return err
	}
	value.Id = id
	value.IsValid = true
	reservationDetails[pos] = value
	return nil
}

func DeleteReservation(id int) int {
	log.Println("Deleting id: ", id)
	for i, v := range reservationDetails {
		if v.Id == id {
			//Found id, delete it
			reservationDetails = append(reservationDetails[:i], reservationDetails[i+1:]...)
			return i
		}
	}
	return -1
}
