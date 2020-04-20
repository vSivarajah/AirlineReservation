package passenger

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Passenger struct {
	Firstname      string `json:firstname`
	Lastname       string `json:lastname`
	Passportnumber int    `json:passportnumber`
	Dateofbirth    string `json:dateofbirth`
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
