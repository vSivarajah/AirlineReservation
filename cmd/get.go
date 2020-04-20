package cmd

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"

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
		Run:   GetReservationById,
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

func GetReservationById(cmd *cobra.Command, args []string) {

	id := cmd.Flag("id")
	idValue, err := strconv.Atoi(id.Value.String())
	url := fmt.Sprintf("http://127.0.0.1:8081/reservation/%d", idValue)
	fmt.Println(url)
	response, err := http.Get(url)
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
