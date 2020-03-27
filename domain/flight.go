package domain

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
