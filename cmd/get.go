package cmd

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/vsivarajah/AirlineReservation/domain/reservations"

	"github.com/spf13/cobra"
	"github.com/vsivarajah/AirlineReservation/cmd/flights"
)

var (
	getCmd = &cobra.Command{
		Use:   "get",
		Short: "Retrivies flight information",
	}

	getFlights = &cobra.Command{
		Use:   "flights",
		Short: "Retrieves all flights",
		Run:   GetFlights,
	}

	getReservation = &cobra.Command{
		Use:   "reservation",
		Short: "Retrives a reservation for the given id",
		Run:   GetReservation,
	}
)

func init() {
	rootCmd.AddCommand(getCmd)
	getCmd.AddCommand(getFlights)
	getCmd.AddCommand(getReservation)
	getReservation.Flags().Int("id", 0, "Specify id to retrieve your booking")

}

func GetFlights(cmd *cobra.Command, args []string) {
	flights.FlightDetails()

}

func GetReservation(cmd *cobra.Command, args []string) {
	id := cmd.Flag("id")
	idValue, err := strconv.Atoi(id.Value.String())
	if err != nil {
		log.Fatal("Could not convert to int", err)
	}
	reservation := GetReservationById(idValue)
	fmt.Println(reservation)
}

func GetReservationById(id int) reservations.Reservation {

	url := fmt.Sprintf("http://127.0.0.1:8081/reservation/%d", id)
	fmt.Println(url)
	response, err := http.Get(url)
	if err != nil {
		log.Fatal(err)
	}
	defer response.Body.Close()
	reservation := reservations.Reservation{}
	err = json.NewDecoder(response.Body).Decode(&reservation)
	if err != nil {
		panic(err)
	}
	return reservation
}
