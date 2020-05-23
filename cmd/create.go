/*
Copyright © 2020 NAME HERE <EMAIL ADDRESS>

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/
package cmd

import (
	"bytes"
	"encoding/json"
	"log"
	"net/http"
	"strconv"

	"github.com/vsivarajah/AirlineReservation/cmd/payments"

	"github.com/spf13/cobra"
	"github.com/vsivarajah/AirlineReservation/cmd/flights"
	"github.com/vsivarajah/AirlineReservation/cmd/passenger"
	"github.com/vsivarajah/AirlineReservation/cmd/reservations"
)

// createCmd represents the create command
var (
	createCmd = &cobra.Command{
		Use:   "create",
		Short: "Creates a reservation, payment.",
	}

	createReservation = &cobra.Command{
		Use:   "reservation",
		Short: "Creates a reservation",
		Run:   CreateReservation,
	}

	createPayment = &cobra.Command{
		Use:   "payment",
		Short: "Makes a payment for given reservation",
		Run:   CreatePayment,
	}
)

func init() {
	rootCmd.AddCommand(createCmd)
	createCmd.AddCommand(createReservation)
	createCmd.AddCommand(createPayment)
	createPayment.Flags().Int("id", 0, "Specify id to retrieve your booking")

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// createCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// createCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

func CreateReservation(cmd *cobra.Command, args []string) {
	traveller := passenger.PassengerDetails()
	sourceairport := "Oslo"
	targetairport := "Cancun"
	flightDetails := flights.FlightInfo{
		SourceAirport: sourceairport,
		TargetAirport: targetairport,
	}

	//flights := FlightInfo{"BOEING777", "Emirates", "Oslo", "Cancun", 2, 2}
	reservation := reservations.ReservationCmd{traveller, flightDetails}
	booking, _ := json.Marshal(reservation)
	req, err := http.NewRequest("POST", "http://127.0.0.1:8081/create", bytes.NewBuffer(booking))
	if err != nil {
		log.Fatal(err)
	}
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()

}

func CreatePayment(cmd *cobra.Command, args []string) {

	id := cmd.Flag("id")
	idValue, err := strconv.Atoi(id.Value.String())
	if err != nil {
		log.Fatal("Could not convert to int", err)
	}

	//reservation := reservations.GetReservationById(idValue)

	payment := payments.Payment{idValue}
	makePayment, _ := json.Marshal(payment)
	req, err := http.NewRequest("POST", "http://127.0.0.1:8081/payment/pay", bytes.NewBuffer(makePayment))
	if err != nil {
		log.Println("tester, går ikke inn her!")
		log.Println(err)
		return
	}
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		log.Println(err)
		return
	}
	defer resp.Body.Close()

	err = reservations.UpdateReservation(idValue)
	if err != nil {
		log.Println("Could not update the reservation")
		return
	}
}
