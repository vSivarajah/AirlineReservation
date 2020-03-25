package domain

type FlightDetail struct {
	Id                int64  `json:"id"`
	SourceAirport     string `json:"sourceairport"`
	TargetAirport     string `json:"targetairport"`
	NumberOfPassenger int64  `json:"numberofpassenger"`
}

type FlightDetails []*FlightDetail

var details = []*FlightDetail{
	&FlightDetail{
		Id:                1,
		SourceAirport:     "Gardermoen",
		TargetAirport:     "India",
		NumberOfPassenger: 2,
	},
}

func CreateFlightDetails(flightDetails *FlightDetail) {
	details = append(details, flightDetails)
}
func GetFlightDetails() FlightDetails {
	return details
}
