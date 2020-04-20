package flights

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
)

type FlightInfo struct {
	SourceAirport string `json:"sourceairport"`
	TargetAirport string `json:"targetairport"`
}

func FlightDetails() []FlightInfo {
	response, err := http.Get("http://127.0.0.1:8081/flights")
	if err != nil {
		log.Fatal(err)
	}
	defer response.Body.Close()

	responseData, err := ioutil.ReadAll(response.Body)
	if err != nil {
		log.Fatal(err)
	}

	var data []FlightInfo

	err = json.Unmarshal([]byte(responseData), &data)
	if err != nil {
		log.Fatal(err)
	}
	return data
}
