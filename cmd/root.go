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
	"bufio"
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/vsivarajah/AirlineReservation/app"

	"github.com/spf13/cobra"

	homedir "github.com/mitchellh/go-homedir"
	"github.com/spf13/viper"
)

var cfgFile string

type ReservationCmd struct {
	Passenger  Passenger  `json:"passenger"`
	FlightInfo FlightInfo `json:"flightinfo"`
}

type Passenger struct {
	Firstname      string `json:firstname`
	Lastname       string `json:lastname`
	Passportnumber int    `json:passportnumber`
	Dateofbirth    string `json:dateofbirth`
}

type FlightInfo struct {
	FlightNumber      string `json:"flightnumber"`
	OperatingAirlines string `json:"operatingairlines"`
	SourceAirport     string `json:"sourceairport"`
	TargetAirport     string `json:"targetairport"`
	MaxSeats          int    `json:"maxseats"`
	NumSeats          int    `json:"numseats"`
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
	rootCmd.AddCommand(PrintTimeCmd)
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

var PrintTimeCmd = &cobra.Command{
	Use: "erik",
	Run: func(cmd *cobra.Command, args []string) {
		now := time.Now()
		prettyTime := now.Format(time.RubyDate)
		println("Hey, the current time is ", prettyTime)
	},
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
		flightDetails := FlightDetails()
		fmt.Println(flightDetails)
	},
}

var CreateReservation = &cobra.Command{
	Use:   "create",
	Short: "Print anything to the screen",
	Long: `print is for printing anything back to the screen.
  For many years people have printed back to the screen.`,
	Run: func(cmd *cobra.Command, args []string) {
		traveller := PassengerDetails()
		flights := FlightDetails()

		//flights := FlightInfo{"BOEING777", "Emirates", "Oslo", "Cancun", 2, 2}
		reservation := ReservationCmd{traveller, flights[0]}
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

func PassengerDetails() Passenger {
	passengerDetails := Passenger{}
	fmt.Println("Please enter your firstname: ")
	firstname_input := bufio.NewReader(os.Stdin)
	firstname, err := firstname_input.ReadString('\n')
	firstname = strings.TrimSuffix(firstname, "\n")
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
	}
	passengerDetails.Firstname = firstname

	fmt.Println("Please enter your lastname: ")
	lastname_input := bufio.NewReader(os.Stdin)
	lastname, err := lastname_input.ReadString('\n')
	lastname = strings.TrimSuffix(lastname, "\n")
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
	}
	passengerDetails.Lastname = lastname

	fmt.Println("Please enter your passportnumber: ")
	passportnumber_input := bufio.NewReader(os.Stdin)
	passportnumber, err := passportnumber_input.ReadString('\n')
	passportnumber = strings.TrimSuffix(passportnumber, "\n")
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
	}
	passportnumberInt, _ := strconv.Atoi(passportnumber)
	passengerDetails.Passportnumber = passportnumberInt

	fmt.Println("Please enter your dateofbirth: ")
	date_input := bufio.NewReader(os.Stdin)
	dateofbirth, err := date_input.ReadString('\n')
	dateofbirth = strings.TrimSuffix(dateofbirth, "\n")
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
	}
	passengerDetails.Dateofbirth = dateofbirth
	return passengerDetails
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
