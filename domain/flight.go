package domain

type Flight struct {
	FlightNumber      string `json:"flightnumber"`
	OperatingAirlines string `json:"operatingairlines"`
	SourceAirport     string `json:"sourceairport"`
	TargetAirport     string `json:"targetairport"`
}
