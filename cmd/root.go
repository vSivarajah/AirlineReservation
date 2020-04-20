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
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/vsivarajah/AirlineReservation/app"
	"github.com/vsivarajah/AirlineReservation/cmd/flights"
	"github.com/vsivarajah/AirlineReservation/cmd/passenger"

	"github.com/spf13/cobra"

	homedir "github.com/mitchellh/go-homedir"
	"github.com/spf13/viper"
)

var cfgFile string

type ReservationCmd struct {
	Passenger  passenger.Passenger `json:"passenger"`
	FlightInfo flights.FlightInfo  `json:"flightinfo"`
}

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "AirlineReservation",
	Short: "A brief description of your application",
	Long: `A longer description that spans multiple lines and likely contains
examples and usage of using your application. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	// Uncomment the following line if your bare application
	// has an action associated with it:
	//	Run: func(cmd *cobra.Command, args []string) { },
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func init() {
	cobra.OnInitialize(initConfig)

	// Here you will define your flags and configuration settings.
	// Cobra supports persistent flags, which, if defined here,
	// will be global for your application.

	rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.AirlineReservation.yaml)")

	// Cobra also supports local flags, which will only run
	// when this action is called directly.
	rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
	rootCmd.AddCommand(StartServer)
	rootCmd.AddCommand(GetFlights)
	rootCmd.AddCommand(CreateReservation)
}

// initConfig reads in config file and ENV variables if set.
func initConfig() {
	if cfgFile != "" {
		// Use config file from the flag.
		viper.SetConfigFile(cfgFile)
	} else {
		// Find home directory.
		home, err := homedir.Dir()
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}

		// Search config in home directory with name ".AirlineReservation" (without extension).
		viper.AddConfigPath(home)
		viper.SetConfigName(".AirlineReservation")
	}

	viper.AutomaticEnv() // read in environment variables that match

	// If a config file is found, read it in.
	if err := viper.ReadInConfig(); err == nil {
		fmt.Println("Using config file:", viper.ConfigFileUsed())
	}
}

var StartServer = &cobra.Command{
	Use: "start",
	Run: func(cmd *cobra.Command, args []string) {
		app.StartApplication()
	},
}

var GetFlights = &cobra.Command{
	Use: "get flights",
	Run: func(cmd *cobra.Command, args []string) {
		flightDetails := flights.FlightDetails()
		fmt.Println(flightDetails)
	},
}

var CreateReservation = &cobra.Command{
	Use:   "create",
	Short: "Print anything to the screen",
	Long: `print is for printing anything back to the screen.
  For many years people have printed back to the screen.`,
	Run: func(cmd *cobra.Command, args []string) {
		traveller := passenger.PassengerDetails()
		sourceairport := "Oslo"
		targetairport := "Cancun"
		flightDetails := flights.FlightInfo{
			SourceAirport: sourceairport,
			TargetAirport: targetairport,
		}

		//flights := FlightInfo{"BOEING777", "Emirates", "Oslo", "Cancun", 2, 2}
		reservation := ReservationCmd{traveller, flightDetails}
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

	},
}
