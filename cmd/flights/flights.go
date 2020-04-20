package flights

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

type FlightInfo struct {
	SourceAirport string `json:"sourceairport"`
	TargetAirport string `json:"targetairport"`
}

func FlightDetails() {
	response, err := http.Get("http://127.0.0.1:8081/flights")
	if err != nil {
		log.Fatal(err)
	}
	defer response.Body.Close()

	responseData, err := ioutil.ReadAll(response.Body)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(string(responseData))
}
