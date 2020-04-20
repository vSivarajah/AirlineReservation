package flights

import (
	"fmt"
	"log"
)

type Flight struct {
	FlightNumber      string `json:"flightnumber"`
	OperatingAirlines string `json:"operatingairlines"`
	SourceAirport     string `json:"sourceairport"`
	TargetAirport     string `json:"targetairport"`
	MaxSeats          int    `json:"maxseats"`
	NumSeats          int    `json:"numseats"`
}

type Flights []*Flight

var ListFlights = []*Flight{
	&Flight{
		FlightNumber:      "BOEING777",
		OperatingAirlines: "Emirates",
		SourceAirport:     "Oslo",
		TargetAirport:     "Cancun",
		MaxSeats:          2,
		NumSeats:          0,
	},
}

// GetFlights return list of flights
func GetFlights() Flights {
	return ListFlights
}

func DoesFlightExist(sourceairport string, targetairport string) bool {
	flights := GetFlights()
	for _, flight := range flights {
		if flight.SourceAirport == sourceairport && flight.TargetAirport == targetairport {
			return true
		}
	}
	return false
}

// AssignFlightNumber assigns flight to reservation
func AssignFlightNumber(sourceairport string, targetairport string) (string, string, int, int) {
	flights := GetFlights()
	for _, flight := range flights {
		if flight.SourceAirport == sourceairport && flight.TargetAirport == targetairport {
			if flight.NumSeats < flight.MaxSeats {
				flight.NumSeats++
				return flight.FlightNumber, flight.OperatingAirlines, flight.MaxSeats, flight.NumSeats
			} else {
				log.Println("Flight is full")
				return "full", "", 0, 0
			}
		}
	}
	value, err := fmt.Println("Can not assign flightnumber as it does not exist")
	if err != nil {
		log.Fatal(err)
	}
	return "", string(value), 0, 0
}

// RemoveReservationFromFlight removes person from flight
func RemoveReservationFromFlight(flightnumber string) bool {
	flights := GetFlights()
	for _, flight := range flights {
		if flight.FlightNumber == flightnumber {
			flight.NumSeats--
			return true
		}
	}
	return false
}
