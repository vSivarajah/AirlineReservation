package cmd

import (
	"fmt"
	"log"
	"strconv"

	"github.com/spf13/cobra"
	"github.com/vsivarajah/AirlineReservation/cmd/flights"
	"github.com/vsivarajah/AirlineReservation/cmd/reservations"
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
	reservation, err := reservations.GetReservationById(idValue)
	if err != nil {
		log.Fatal("Could not fetch the reservation", err)
	}
	fmt.Println(reservation)
}
