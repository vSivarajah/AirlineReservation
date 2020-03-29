package domain

import (
	"fmt"
	"log"
)

type Flight struct {
	FlightNumber      string `json:"flightnumber"`
	OperatingAirlines string `json:"operatingairlines"`
	SourceAirport     string `json:"sourceairport"`
	TargetAirport     string `json:"targetairport"`
}

type Flights []*Flight

var ListFlights = []*Flight{
	&Flight{
		FlightNumber:      "BOEING777",
		OperatingAirlines: "Emirates",
		SourceAirport:     "Oslo",
		TargetAirport:     "Cancun",
	},
}

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
func AssignFlightNumber(sourceairport string, targetairport string) (string, string) {
	flights := GetFlights()
	for _, flight := range flights {
		if flight.SourceAirport == sourceairport && flight.TargetAirport == targetairport {
			return flight.FlightNumber, flight.OperatingAirlines
		}
	}
	value, err := fmt.Println("Can not assign flightnumber as it does not exist")
	if err != nil {
		log.Fatal(err)
	}
	return "", string(value)
}
