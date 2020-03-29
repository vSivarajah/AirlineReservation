package services

import "github.com/vsivarajah/AirlineReservation/domain/flights"

var (
	FlightService flightServiceInterface = &flightService{}
)

type flightService struct{}

type flightServiceInterface interface {
	AssignFlightNumber(string, string) (string, string)
	DoesFlightExist(string, string) bool
	GetFlights() flights.Flights
}

func (s *flightService) DoesFlightExist(sourceairport string, targetairport string) bool {
	return flights.DoesFlightExist(sourceairport, targetairport)
}

func (s *flightService) AssignFlightNumber(sourceairport string, targetairport string) (string, string) {
	return flights.AssignFlightNumber(sourceairport, targetairport)
}

func (s *flightService) GetFlights() flights.Flights {
	return flights.GetFlights()
}
